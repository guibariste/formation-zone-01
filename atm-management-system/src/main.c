#include "header.h"

void mainMenu(struct User u)
{
    int option;
    int account_number;
    system("clear");
    printf("\n\n\t\t======= ATM =======\n\n");
    printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
    printf("\n\t\t[1]- Create a new account\n");
    printf("\n\t\t[2]- Update account information\n");
    printf("\n\t\t[3]- Check accounts\n");
    printf("\n\t\t[4]- Check list of owned account\n");
    printf("\n\t\t[5]- Make Transaction\n");
    printf("\n\t\t[6]- Remove existing account\n");
    printf("\n\t\t[7]- Transfer ownership\n");
    printf("\n\t\t[8]- Exit\n");
    scanf("%d", &option);

    switch (option)
    {
    case 1:
        createNewAcc(u);
        break;
    case 2:
        // student TODO : add your **Update account information** function
        ModifyAccount(u);
        // here
        break;
    case 3:
        // student TODO : add your **Check the details of existing accounts** function
        CheckingDetailAccount(u);
        // here
        break;
    case 4:
        checkAllAccounts(u);
        break;
    case 5:
        // student TODO : add your **Make transaction** function
        MakeTransaction(u);
        // here
        break;
    case 6:
        
        // student TODO : add your **Remove existing account** function
        // here
        
        printf("\nEnter the Account Number:");
        scanf("%d", &account_number);
        RemoveAccount(u, account_number);
        //RemoveExistingUser(u);
        break;
    case 7:
        // student TODO : add your **Transfer owner** function
        // here
        TransferOwner(u);
        break;
    case 8:
        exit(1);
        break;
    default:
        printf("Invalid operation!\n");
    }
};

void initMenu(struct User *u)
{
    int r = 0;
    int option;
    system("clear");
    printf("\n\n\t\t======= ATM =======\n");
    printf("\n\t\t-->> Feel free to login / register :\n");
    printf("\n\t\t[1]- login\n");
    printf("\n\t\t[2]- register\n");
    printf("\n\t\t[3]- exit\n");
    while (!r)
    {
        scanf("%d", &option);
        switch (option)
        {
        case 1:
            loginMenu(u->name, u->password);
            //printf("%s\n",u->password);
            if (strcmp(u->password, getPassword(*u)) == 0)
            {
                printf("\n\nPassword Match!");
            }
            else
            {
                printf("\nWrong password!! or User Name\n");
                exit(1);
            }
            r = 1;
            break;
        case 2:
            // student TODO : add your **Registration** function
            RegistrationMenu(u->name, u->password);
           
            // here
            if ((strcmp(u->name, getUser(*u))==0))
            {
                printf("\n\nCette utilisateur est deja inscrit\n\n");
                exit(1);
            }
            else
            {
                printf("\n\nwait");
                AddNewUser(*u);
                
            }
            r = 1;
            break;
        case 3:
            exit(1);
            break;
        default:
            printf("Insert a valid operation!\n");
        }
    }
};


void AddNewUser(struct User u)
{
    FILE *fileptr1, *fileptr2;
    struct User userChecker;
    char filename[40]="./data/users.txt";
    char ch;
    int count=0;
    int temp=0;
    if ((fileptr1 = fopen(filename, "a+")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    else
    {   
        char count_char[10];
        char len;
        count=AutoIncrementID();
        //fprintf(fileptr1,"\n");
        sprintf(count_char, "%d", count);  // convert int to string
        fprintf(fileptr1, "%s %s %s\n",count_char,u.name, u.password); // write string to file
        fclose(fileptr1);
    }
    
    //remove(filename);
    //rename the file replica.c to original name
    //rename("./data/replica.c", filename);
    
}

int main()
{
    struct User u;
    
    initMenu(&u);
    mainMenu(u);
    return 0;
}
