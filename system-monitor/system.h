#ifndef SYSTEM_H
#define SYSTEM_H

#include <string>


std::string CPUinfo();
const char* getOsName();
std ::string getLoggedInUsername();
std::string getHostName();
int getProcessCount();

#endif // SYSTEM_H
