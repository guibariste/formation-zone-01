#ifndef MEM_H
#define MEM_H

#include <string>

struct MemoryInfo
{
    std::string used;
    std::string total;
    std::string swap;
};
struct DiskInfo
{
    std::string total;
    std::string used;
};


MemoryInfo getMemoryUsage();
DiskInfo getDiskUsage();


#endif // SYSTEM_H