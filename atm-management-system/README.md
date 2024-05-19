# atm-system

## Functions
``` C

// authentication functions
void loginMenu(char a[50], char pass[50]);
void RegistrationMenu(char a[50], char pass[50]);
const char *getPassword(struct User u);
const char *getUser(struct User u);

// system function
void createNewAcc(struct User u);
void mainMenu(struct User u);
void checkAllAccounts(struct User u);
void RemoveExistingUser(struct User u);
void AddNewUser(struct User u);
void RemoveAccount(struct User u, int account_number);
void CheckingDetailAccount(struct User u);
void ModifyAccount(struct User u);
void TransferOwner(struct User u);
int CheckIfUserExist(struct User u, char userName[50]);
int GetID(struct User u);

```
### Installation
First Install gcc and make
```
sudo apt-get install gcc make
```
```
    git clone https://zone01normandie.org/git/maximediet/atm-system.git
    cd atm-system
    make
    ./atm
```