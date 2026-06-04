# Unit Testing

## Overview

The project uses **testify** for assertions and **mockery** for mock generation. Tests focus on the subsystem layer,
which is the most testable part of the codebase: subsystem functions already receive `*userdata.Session` and
`*gamedata.Gamedata` as explicit parameters, making them testable with the right helpers in place.

---

## Frameworks

| Library                        | Purpose                                         |
|--------------------------------|-------------------------------------------------|
| `github.com/stretchr/testify`  | Assertions (`assert`/`require`) and test suites |
| `github.com/vektra/mockery/v2` | Auto-generates typed mocks from Go interfaces   |

Install:

```bash
go get github.com/stretchr/testify@latest
go get github.com/vektra/mockery/v2@latest
```

---

## Test Infrastructure

Four helper packages reduce boilerplate across all tests. They live alongside the packages they support and are only
compiled when running tests.

### 1. Config helper — `internal/config/testhelpers.go`

`config.Conf` is a mutable global. The helper provides a safe default and a standard save/restore pattern so tests do
not interfere with each other.

```go
// Returns a RuntimeConfig with sensible test defaults.
func TestConf() *RuntimeConfig

// Standard usage in any test that touches config:
orig := config.Conf
defer func () { config.Conf = orig }()
config.Conf = config.TestConf()
```

### 2. Gamedata builder — `internal/gamedata/testbuilder.go`

`*gamedata.Gamedata` has 50+ map fields. The builder constructs a minimal, in-memory Gamedata without touching any
database file. Add `With*` methods for each field as new subsystems are covered.

```go
gd := gamedata.NewTestGamedata().
WithGachaGroup(100, &gamedata.GachaGroup{GroupMasterId: 100, GroupWeight: 85}).
WithGacha(1, &gamedata.Gacha{GachaMasterId: 1, GachaGroups: []int32{100}}).
Build()
```

### 3. Session factory — `internal/testutil/session.go`

Creates a `*userdata.Session` with only the fields populated that the function under test actually reads. This avoids
touching the database engine or Gin context.

```go
// Session with gamedata and UserStatus populated, no DB.
func NewTestSession(gd *gamedata.Gamedata, status *client.UserStatus) *userdata.Session

// Bare session for pure functions that only need gamedata maps.
func NewMinimalSession(gd *gamedata.Gamedata) *userdata.Session
```

### 4. Database interface + mock — `internal/userdata/db_interface.go`

Most subsystems call `session.Db` (a `*xorm.Session`) directly. A thin interface over the methods actually used enables
mocking without a real database. Only the methods needed by the subsystems currently under test are required.

```go
// internal/userdata/db_interface.go
type UserDb interface {
Table(table string) UserDb
Where(query interface{}, args ...interface{}) UserDb
Find(dest interface{}) error
Insert(dest ...interface{}) (int64, error)
Update(dest interface{}) (int64, error)
Delete(dest interface{}) (int64, error)
}
```

The concrete implementation wraps `*xorm.Session` (`internal/userdata/xorm_db.go`). Mockery generates the mock
automatically — see [Mock Generation](#mock-generation) below.

---

## Writing a Test

### Pure functions (no DB, no session)

Some subsystem functions are completely stateless. Test them directly:

```go
// internal/subsystem/user_gacha/gacha_test.go
package user_gacha

import (
	"testing"

	"elichika/internal/gamedata"
	"elichika/internal/serverdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChooseRandomCard_SingleCard(t *testing.T) {
	gd := gamedata.NewTestGamedata().
		WithGachaGroup(1, &gamedata.GachaGroup{GroupMasterId: 1, GroupWeight: 100}).
		Build()

	cards := []serverdata.GachaCard{{GroupMasterId: 1, CardMasterId: 42}}

	result := ChooseRandomCard(gd, cards)
	assert.Equal(t, int32(42), result)
}

func TestChooseRandomCard_EmptyPool(t *testing.T) {
	gd := gamedata.NewTestGamedata().Build()
	result := ChooseRandomCard(gd, nil)
	assert.Equal(t, int32(0), result)
}
```

### Subsystem functions that read session state (no DB writes)

Functions that read `session.Gamedata` or `session.UserStatus` but don't write to the database:

```go
// internal/subsystem/user_gacha/get_gacha_list_test.go
package user_gacha

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/gamedata"
	"elichika/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetGachaList_ExcludesTutorialGachaAfterTutorial(t *testing.T) {
	tutorialGacha := &gamedata.Gacha{GachaMasterId: 999999}
	normalGacha := &gamedata.Gacha{GachaMasterId: 1}

	gd := gamedata.NewTestGamedata().
		WithGachaInList(normalGacha).
		WithGachaInList(tutorialGacha).
		Build()

	status := &client.UserStatus{TutorialPhase: enum.TutorialPhaseFinal}
	session := testutil.NewMinimalSession(gd)
	session.UserStatus = status

	result := GetGachaList(session)

	assert.Equal(t, 1, result.Size())
	assert.Equal(t, int32(1), result.Get(0).GachaMasterId)
}

func TestGetGachaList_IncludesTutorialGachaDuringTutorial(t *testing.T) {
	tutorialGacha := &gamedata.Gacha{GachaMasterId: 999999}

	gd := gamedata.NewTestGamedata().
		WithGachaInList(tutorialGacha).
		Build()

	status := &client.UserStatus{TutorialPhase: enum.TutorialPhaseGacha}
	session := testutil.NewMinimalSession(gd)
	session.UserStatus = status

	result := GetGachaList(session)

	assert.Equal(t, 1, result.Size())
}
```

### Subsystem functions that write to the database

Use the mockery-generated `UserDb` mock to set expectations on which tables and queries are called:

```go
// internal/subsystem/user_live/clear_user_live_test.go
package user_live

import (
	"testing"

	"elichika/internal/testutil"
	"elichika/internal/testutil/mocks"
	"github.com/stretchr/testify/mock"
)

func TestClearUserLive(t *testing.T) {
	db := mocks.NewUserDb(t)
	db.EXPECT().Table("u_live").Return(db)
	db.EXPECT().Where("user_id = ?", int32(42)).Return(db)
	db.EXPECT().Delete(mock.Anything).Return(int64(1), nil)

	session := testutil.NewTestSession(nil, nil)
	session.UserId = 42
	session.Db = db // inject mock

	ClearUserLive(session)

	db.AssertExpectations(t)
}
```

---

## Mock Generation

Mockery is configured via `.mockery.yaml` at the project root:

```yaml
with-expecter: true
packages:
  elichika/internal/userdata:
    interfaces:
      UserDb:
        config:
          dir: internal/testutil/mocks
          outpkg: mocks
```

Regenerate mocks after changing the `UserDb` interface:

```bash
go run github.com/vektra/mockery/v2
```

The generated file lands at `internal/testutil/mocks/UserDb.go` and should be committed alongside the interface change.

---

## File Layout

```
internal/
├── config/
│   └── testhelpers.go          ← TestConf() and save/restore pattern
├── gamedata/
│   └── testbuilder.go          ← NewTestGamedata() builder
├── testutil/
│   ├── session.go              ← NewTestSession(), NewMinimalSession()
│   └── mocks/
│       └── UserDb.go           ← mockery-generated mock (committed)
├── userdata/
│   ├── db_interface.go         ← UserDb interface
│   └── xorm_db.go              ← concrete xorm wrapper
└── subsystem/
    └── user_gacha/
        ├── gacha_test.go       ← pilot: ChooseRandomCard tests
        └── get_gacha_list_test.go  ← pilot: GetGachaList tests
.mockery.yaml                   ← mockery configuration
```

---

## Running Tests

```bash
# Run all tests
go test ./internal/...

# Run a specific subsystem
go test ./internal/subsystem/user_gacha/...

# With coverage report
go test -cover ./internal/subsystem/user_gacha/...

# With race detector
go test -race ./internal/subsystem/user_gacha/...

# Verbose output
go test -v ./internal/subsystem/user_gacha/...
```

---

## Extending to a New Subsystem

Once the infrastructure is in place, adding tests for a new subsystem is a short checklist:

1. Add `With*` methods to `GamedataBuilder` for any new field types that subsystem reads.
2. Add any new method signatures to the `UserDb` interface that the subsystem calls.
3. Re-run `go run github.com/vektra/mockery/v2` to regenerate the mock.
4. Write `<file>_test.go` alongside the subsystem file, following the patterns above.
