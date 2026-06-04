# Account Management

Elichika supports multiple accounts — either multiple accounts for a single person, or multiple different users on the
same server.

---

## Creating an Account

New accounts are created automatically when you log in with a User ID that does not exist in the database. You can also
create accounts from the WebUI's Account Builder.

---

## Account Transfer (Switching Accounts)

Use the in-game transfer system to switch between accounts or create a new account. Select **"Transfer with password"
** — Google Play Game Services is not currently supported.

![Transfer system](images/transfer_1.png)

Enter your User ID and password:

- **User ID** is a non-negative integer of up to 9 digits. For switching accounts, this is your existing Player ID.
- If the User ID exists, the entered password is verified against the stored password.
- If the User ID does not exist, a new account is created with that ID and password. You may leave the password empty,
  but if you are on a shared server, setting a password is strongly recommended — anyone who knows your User ID can
  claim the account without one.
- Passwords are stored securely with bcrypt.

![Enter ID and password](images/transfer_2.png)

Confirm the transfer to log in as the new user.

![Confirm transfer](images/transfer_3.png)

You can change your password at any time using the in-game transfer system:

![Use the transfer system](images/transfer_4.png)
![Set new password](images/transfer_5.png)
![Result](images/transfer_6.png)

---

## Multiple Devices

You can play from multiple devices connected to the same external server, as long as only one device is active at a
time. Logging in from a second device disconnects the first, preventing conflicting writes to your account data.

> This does not apply to the embedded server or a server run inside Termux, which do not support concurrent connections
> from separate devices.

---

## Japanese vs. Global Client

Both the Japanese and Global clients work against the same server and database. However, it is recommended not to play a
single account on both clients simultaneously — some content is exclusive to one region and will cause the client to
freeze if accessed from the wrong one.
