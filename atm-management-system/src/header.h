#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Date
{
    int month, day, year;
};

// all fields for each record of an account
struct Record
{
    int id;
    int userId;
    char name[100];
    char country[100];
    int phone;
    char accountType[10];
    int accountNbr;
    double amount;
    struct Date deposit;
    struct Date withdraw;
};

struct User
{
    int id;
    char name[50];
    char password[50];
};

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
void MakeTransaction(struct User u);
void GetAccountByUser(int account_number, struct Record *result, struct User u);
int CheckIfAccountIsOwnedByUser(struct User u, char userName[50], int accountNbr);
int AutoIncrementID();
