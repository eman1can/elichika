# Login Bonus Information

There are 2 categories of login bonuses:
 - 2d "IllustLoginBonus" with a 2d image background
 - 3d "NaviLoginBonus" with the partner or bday girl, etc.

There are 6 types of login bonuses:
 - Normal: Normal login bonus, the common login bonus that loops
 - Beginner: The beginner login bonus, which only runs for new players
 - Event 2d: The 2d event login bonus, which runs during events and has a 2d image background
 - Event 3d: The 3d event login bonus, which runs during events and has a 3d model background
 - Birthday: A single-day login bonus that triggers on the birthday of a member
 - Comeback: A special login bonus for players who haven't logged in for at least 14 days.

Login Bonus Info is stored in 3 server tables:
- `s_login_bonus`:
  - Stores id, background and whiteboard ids, start and end time. Only id, start and end are required
- `s_login_bonus_reward_day`
  - Stores the day of the login bonus and the content grade of the reward
- `s_login_bonus_reward_content` 
  - Stores the content of the login bonus days

User info for the login bonus is stored in `u_login_bonus`:
- `user_id`: The ID of the user
- `login_bonus_id`: The ID of the login bonus
- `last_received_reward`: the 0-indexed reward, default to -1 (have not received anything)
- `last_received_at`: timestamp of the last time rewarded, default to 0

Login Bonus workflow is:
1. User calls bootstrap to get login bonus
2. Each login bonus is checked and rewarded if eligible
3. Then the client send read requests that are not tracked
4. from recorded data, the present box count go up only after the read requests are sent, but the present count doesn't go up during the read request
5. this can interpreted as the server fill-in the present box data first, and then add present for the login bonus on top of that
6. although it can also be interpreted as the server only award item once read is called.
