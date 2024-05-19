#include "header.h"
#include <pwd.h>
#include <unistd.h>
#include <fstream>
#include <string>
#include <iostream>


#include <sstream>


#include <cstdio>









#include <cstdlib>


#include <dirent.h>
#include <cctype>

bool isNumeric(const char* str)
{
    for (int i = 0; str[i] != '\0'; i++)
    {
        if (!std::isdigit(str[i]))
        {
            return false;
        }
    }
    return true;
}

int getProcessCount()
{
    int count = 0;
    DIR* procDir = opendir("/proc");
    if (procDir)
    {
        struct dirent* entry;
        while ((entry = readdir(procDir)) != NULL)
        {
            // Check if the entry is a directory and its name is a number
            if (entry->d_type == DT_DIR && isNumeric(entry->d_name))
            {
                count++;
            }
        }
        closedir(procDir);
    }
    return count;
}


// get cpu id and information, you can use `proc/cpuinfo`
string CPUinfo()
{
    char CPUBrandString[0x40];
    unsigned int CPUInfo[4] = {0, 0, 0, 0};

    // unix system
    // for windoes maybe we must add the following
    // __cpuid(regs, 0);
    // regs is the array of 4 positions
    __cpuid(0x80000000, CPUInfo[0], CPUInfo[1], CPUInfo[2], CPUInfo[3]);
    unsigned int nExIds = CPUInfo[0];

    memset(CPUBrandString, 0, sizeof(CPUBrandString));

    for (unsigned int i = 0x80000000; i <= nExIds; ++i)
    {
        __cpuid(i, CPUInfo[0], CPUInfo[1], CPUInfo[2], CPUInfo[3]);

        if (i == 0x80000002)
            memcpy(CPUBrandString, CPUInfo, sizeof(CPUInfo));
        else if (i == 0x80000003)
            memcpy(CPUBrandString + 16, CPUInfo, sizeof(CPUInfo));
        else if (i == 0x80000004)
            memcpy(CPUBrandString + 32, CPUInfo, sizeof(CPUInfo));
    }
    string str(CPUBrandString);
    return str;
}

// getOsName, this will get the OS of the current computer
const char *getOsName()
{
#ifdef _WIN32
    return "Windows 32-bit";
#elif _WIN64
    return "Windows 64-bit";
#elif __APPLE__ || __MACH__
    return "Mac OSX";
#elif __linux__
    return "Linux";
#elif __FreeBSD__
    return "FreeBSD";
#elif __unix || __unix__
    return "Unix";
#else
    return "Other";
#endif
}
// getLoggedInUsername, this will get the username of the logged-in user
string getLoggedInUsername()
{
    string username;

#ifdef _WIN32
    // Windows implementation
    char buffer[UNLEN + 1];
    DWORD size = sizeof(buffer);
    if (GetUserNameA(buffer, &size))
    {
        username = buffer;
    }

#elif __linux__
    // Linux implementation
    struct passwd *pwd;
    pwd = getpwuid(getuid());
    if (pwd)
    {
        username = pwd->pw_name;
    }

#elif __APPLE__ || __MACH__
    // macOS implementation
    char buffer[256];
    uint32_t size = sizeof(buffer);
    if (getlogin_r(buffer, size) == 0)
    {
        username = buffer;
    }std::string getHostName()
{
    char hostname[256];
    if (gethostname(hostname, sizeof(hostname)) != 0)
    {
        return "Unknown";
    }
    return hostname;
}

#else
    // Unsupported platform
    username = "Unknown";
#endif

    return username;
}
std::string getHostName()
{
    char hostname[256];
    if (gethostname(hostname, sizeof(hostname)) != 0)
    {
        return "Unknown";
    }
    return hostname;
}
