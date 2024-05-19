#include "header.h"
#include <iostream>
#include <sstream>
#include <fstream>
#include <vector>

struct MemoryInfo
{
    std::string total;
    std::string used;
    std::string swap;

};
struct DiskInfo
{
    std::string total;
    std::string used;
};
MemoryInfo getMemoryUsage()
{
    MemoryInfo memoryInfo;

    std::string command = "free -h";
    std::vector<std::string> outputLines;

    FILE* pipe = popen(command.c_str(), "r");
    if (!pipe)
    {
        return memoryInfo;
    }

    char buffer[128];
    while (!feof(pipe))
    {
        if (fgets(buffer, 128, pipe) != NULL)
        {
            outputLines.push_back(buffer);
        }
    }

    pclose(pipe);

    bool memFound = false;
    for (size_t i = 0; i < outputLines.size(); i++)
    {
        std::istringstream iss(outputLines[i]);
        std::string word;

        if (iss >> word && word == "Mem:")
        {
            memFound = true;
            iss >> memoryInfo.used;
            iss >> memoryInfo.total;

            // Get the line following "Mem:"
            if (i + 1 < outputLines.size())
            {
                std::istringstream lineIss(outputLines[i + 1]);
                lineIss >> word; // Skip the first word
                lineIss >> word;
                lineIss >> memoryInfo.swap;
            }

            break;
        }
    }

    return memoryInfo;
}


DiskInfo getDiskUsage()
{
    DiskInfo disk;

    std::string command = "df -h /";
    std::vector<std::string> outputLines;

    FILE* pipe = popen(command.c_str(), "r");
    if (!pipe)
    {
        return disk;
    }

    char buffer[128];
    while (!feof(pipe))
    {
        if (fgets(buffer, 128, pipe) != NULL)
        {
            outputLines.push_back(buffer);
        }
    }

    pclose(pipe);

    if (outputLines.size() >= 2)
    {
        std::istringstream lineIss(outputLines[1]);
        std::string word;

        // Ignorer le premier mot
        lineIss >> word;

        // Récupérer les mots total et used
        lineIss >> disk.total;
        lineIss >> disk.used;
    }

    return disk;
}