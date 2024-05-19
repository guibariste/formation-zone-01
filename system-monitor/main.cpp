#include "header.h"
#include <SDL.h>
#include "system.h"
 #include "mem.h"
#include <regex>


#include "imgui.h"
/*
NOTE : You are free to change the code as you wish, the main objective is to make the
       application work and pass the audit.

       It will be provided the main function with the following functions :

       - `void systemWindow(const char *id, ImVec2 size, ImVec2 position)`
            This function will draw the system window on your screen
       - `void memoryProcessesWindow(const char *id, ImVec2 size, ImVec2 position)`
            This function will draw the memory and processes window on your screen
       - `void networkWindow(const char *id, ImVec2 size, ImVec2 position)`
            This function will draw the network window on your screen
*/

// About Desktop OpenGL function loaders:
//  Modern desktop OpenGL doesn't have a standard portable header file to load OpenGL function pointers.
//  Helper libraries are often used for this purpose! Here we are supporting a few common ones (gl3w, glew, glad).
//  You may use another loader/header of your choice (glext, glLoadGen, etc.), or chose to manually implement your own.
#if defined(IMGUI_IMPL_OPENGL_LOADER_GL3W)
#include <GL/gl3w.h> // Initialize with gl3wInit()
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLEW)
#include <GL/glew.h> // Initialize with glewInit()
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLAD)
#include <glad/glad.h> // Initialize with gladLoadGL()
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLAD2)
#include <glad/gl.h> // Initialize with gladLoadGL(...) or gladLoaderLoadGL()
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLBINDING2)
#define GLFW_INCLUDE_NONE      // GLFW including OpenGL headers causes ambiguity or multiple definition errors.
#include <glbinding/Binding.h> // Initialize with glbinding::Binding::initialize()
#include <glbinding/gl/gl.h>
using namespace gl;
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLBINDING3)
#define GLFW_INCLUDE_NONE        // GLFW including OpenGL headers causes ambiguity or multiple definition errors.
#include <glbinding/glbinding.h> // Initialize with glbinding::initialize()
#include <glbinding/gl/gl.h>
using namespace gl;
#else
#include IMGUI_IMPL_OPENGL_LOADER_CUSTOM
#endif
double calculatePercentage(double used, double total)
{
    if (total == 0)
    {
        return 0.0;
    }

    return (used / total) * 100.0;
}

double extractDigits(const std::string& str)
{
    std::regex digitsRegex("\\d+\\.?\\d*");
    std::smatch match;

    if (std::regex_search(str, match, digitsRegex))
    {
        std::istringstream iss(match[0]);
        double value;
        if (iss >> value)
        {
            return value;
        }
    }

    return 0.0;  // Valeur par défaut si aucun chiffre n'est trouvé ou conversion échouée
}

// systemWindow, display information for the system monitorization
void systemWindow(const char *id, ImVec2 size, ImVec2 position)
{
    ImGui::Begin(id);
    ImGui::SetWindowSize(id, size);
    ImGui::SetWindowPos(id, position);

    // student TODO : add code here for the system window
   const char* osName = getOsName();
    ImGui::Text("Operating System: %s", osName);
    string cpuInfo = CPUinfo();

    // Print the CPU information
    ImGui::Text("CPU Information: %s", cpuInfo.c_str());

 string username = getLoggedInUsername();

    // Print the logged-in username
    ImGui::Text("Logged-in User: %s", username.c_str());

    std::string hostname = getHostName();
    ImGui::Text("Hostname: %s", hostname.c_str());
    int totalProcesses = getProcessCount();
ImGui::Text("Total des tâches : %d", totalProcesses);




    ImGui::End();
}

// memoryProcessesWindow, display information for the memory and processes information
void memoryProcessesWindow(const char *id, ImVec2 size, ImVec2 position)
{
    ImGui::Begin(id);
    ImGui::SetWindowSize(id, size);
    ImGui::SetWindowPos(id, position);

  

   MemoryInfo memoryInfo = getMemoryUsage();
    std::string memoryUsed = memoryInfo.used;
    std::string memoryTotal = memoryInfo.total;
    std::string memorySwap = memoryInfo.swap;



 double memoryUsedValue = extractDigits(memoryUsed);
double memoryTotalValue = extractDigits(memoryTotal);
double memoryswapUsedValue = extractDigits(memorySwap);
double memoryswapTotal = 2.0;
double percentage = calculatePercentage(memoryUsedValue, memoryTotalValue);
double percentageSwap = calculatePercentage(memoryswapUsedValue, memoryswapTotal);



ImVec2 progressBarSize(100.0f, 20.0f);

// Calculer la taille de la partie remplie
float filledWidth = static_cast<float>(progressBarSize.x * percentage / 100.0);

// Dessiner le pourcentage en tant que chaîne de caractères

ImGui::Text("RAM : %.2f%%", percentage);
ImGui::SameLine();
// Dessiner la barre de progression
ImGui::ProgressBar(percentage / 100.0, progressBarSize);
ImGui::SameLine();
ImGui::PushStyleColor(ImGuiCol_PlotHistogram, ImVec4(1.0f, 0.5f, 0.0f, 1.0f));
ImGui::Dummy(ImVec2(filledWidth, progressBarSize.y));
ImGui::PopStyleColor();
ImGui::SameLine();
ImGui::Text("  Memory Used :%s/%s", memoryUsed.c_str(),memoryTotal.c_str());



// ImGui::Text("Memory Swap Percentage: %.2f%%", percentageSwap);

float filledWidt = static_cast<float>(progressBarSize.x * percentageSwap / 100.0);

// Dessiner le pourcentage en tant que chaîne de caractères

ImGui::Text("Swap : %.2f%%", percentageSwap);
ImGui::SameLine();
// Dessiner la barre de progression
ImGui::ProgressBar(percentageSwap / 100.0, progressBarSize);
ImGui::SameLine();
ImGui::PushStyleColor(ImGuiCol_PlotHistogram, ImVec4(1.0f, 0.5f, 0.0f, 1.0f));
ImGui::Dummy(ImVec2(filledWidt, progressBarSize.y));
ImGui::PopStyleColor();
ImGui::SameLine();
ImGui::Text("Used swap: %s/2GB", memorySwap.c_str());



DiskInfo diskInfo = getDiskUsage();
    std::string diskUsed = diskInfo.used;
   std::string diskTotal = diskInfo.total;

   double convDiskUsed = extractDigits(diskUsed);
double convDiskTotal = extractDigits(diskTotal);
double percentageDisk = calculatePercentage(convDiskUsed, convDiskTotal);
    
    //  ImGui::Text("Disk total %s", diskTotal.c_str());

float filledWid = static_cast<float>(progressBarSize.x * percentageDisk / 100.0);

// Dessiner le pourcentage en tant que chaîne de caractères

ImGui::Text("Disk : %.2f%%", percentageDisk);
ImGui::SameLine();
// Dessiner la barre de progression
ImGui::ProgressBar(percentageDisk / 100.0, progressBarSize);
ImGui::SameLine();
ImGui::PushStyleColor(ImGuiCol_PlotHistogram, ImVec4(1.0f, 0.5f, 0.0f, 1.0f));
ImGui::Dummy(ImVec2(filledWid, progressBarSize.y));
ImGui::PopStyleColor();
ImGui::SameLine();
ImGui::Text("  Disk used :%s/%s", diskUsed.c_str(),diskTotal.c_str());



                // ImGui::Text("Disk used  %s", diskUsed.c_str());
                // ImGui::Text("Memory Disk Percentage: %.2f%%", percentageDisk);


    ImGui::End();
}

// network, display information network information
void networkWindow(const char *id, ImVec2 size, ImVec2 position)
{
    ImGui::Begin(id);
    ImGui::SetWindowSize(id, size);
    ImGui::SetWindowPos(id, position);

    // student TODO : add code here for the network information

    ImGui::End();
}

// Main code
int main(int, char **)
{
    // Setup SDL
    // (Some versions of SDL before <2.0.10 appears to have performance/stalling issues on a minority of Windows systems,
    // depending on whether SDL_INIT_GAMECONTROLLER is enabled or disabled.. updating to latest version of SDL is recommended!)
    if (SDL_Init(SDL_INIT_VIDEO | SDL_INIT_TIMER | SDL_INIT_GAMECONTROLLER) != 0)
    {
        printf("Error: %s\n", SDL_GetError());
        return -1;
    }

    // GL 3.0 + GLSL 130
    const char *glsl_version = "#version 130";
    SDL_GL_SetAttribute(SDL_GL_CONTEXT_FLAGS, 0);
    SDL_GL_SetAttribute(SDL_GL_CONTEXT_PROFILE_MASK, SDL_GL_CONTEXT_PROFILE_CORE);
    SDL_GL_SetAttribute(SDL_GL_CONTEXT_MAJOR_VERSION, 3);
    SDL_GL_SetAttribute(SDL_GL_CONTEXT_MINOR_VERSION, 0);

    // Create window with graphics context
    SDL_GL_SetAttribute(SDL_GL_DOUBLEBUFFER, 1);
    SDL_GL_SetAttribute(SDL_GL_DEPTH_SIZE, 24);
    SDL_GL_SetAttribute(SDL_GL_STENCIL_SIZE, 8);
    SDL_WindowFlags window_flags = (SDL_WindowFlags)(SDL_WINDOW_OPENGL | SDL_WINDOW_RESIZABLE | SDL_WINDOW_ALLOW_HIGHDPI);
    SDL_Window *window = SDL_CreateWindow("Dear ImGui SDL2+OpenGL3 example", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, 1280, 720, window_flags);
    SDL_GLContext gl_context = SDL_GL_CreateContext(window);
    SDL_GL_MakeCurrent(window, gl_context);
    SDL_GL_SetSwapInterval(1); // Enable vsync

    // Initialize OpenGL loader
#if defined(IMGUI_IMPL_OPENGL_LOADER_GL3W)
    bool err = gl3wInit() != 0;
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLEW)
    bool err = glewInit() != GLEW_OK;
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLAD)
    bool err = gladLoadGL() == 0;
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLAD2)
    bool err = gladLoadGL((GLADloadfunc)SDL_GL_GetProcAddress) == 0; // glad2 recommend using the windowing library loader instead of the (optionally) bundled one.
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLBINDING2)
    bool err = false;
    glbinding::Binding::initialize();
#elif defined(IMGUI_IMPL_OPENGL_LOADER_GLBINDING3)
    bool err = false;
    glbinding::initialize([](const char *name) { return (glbinding::ProcAddress)SDL_GL_GetProcAddress(name); });
#else
    bool err = false; // If you use IMGUI_IMPL_OPENGL_LOADER_CUSTOM, your loader is likely to requires some form of initialization.
#endif
    if (err)
    {
        fprintf(stderr, "Failed to initialize OpenGL loader!\n");
        return 1;
    }

    // Setup Dear ImGui context
    IMGUI_CHECKVERSION();
    ImGui::CreateContext();
    // render bindings
    ImGuiIO &io = ImGui::GetIO();

    // Setup Dear ImGui style
    ImGui::StyleColorsDark();

    // Setup Platform/Renderer backends
    ImGui_ImplSDL2_InitForOpenGL(window, gl_context);
    ImGui_ImplOpenGL3_Init(glsl_version);

    // background color
    // note : you are free to change the style of the application
    ImVec4 clear_color = ImVec4(0.0f, 0.0f, 0.0f, 0.0f);

    // Main loop
    bool done = false;
    while (!done)
    {
        // Poll and handle events (inputs, window resize, etc.)
        // You can read the io.WantCaptureMouse, io.WantCaptureKeyboard flags to tell if dear imgui wants to use your inputs.
        // - When io.WantCaptureMouse is true, do not dispatch mouse input data to your main application.
        // - When io.WantCaptureKeyboard is true, do not dispatch keyboard input data to your main application.
        // Generally you may always pass all inputs to dear imgui, and hide them from your application based on those two flags.
        SDL_Event event;
        while (SDL_PollEvent(&event))
        {
            ImGui_ImplSDL2_ProcessEvent(&event);
            if (event.type == SDL_QUIT)
                done = true;
            if (event.type == SDL_WINDOWEVENT && event.window.event == SDL_WINDOWEVENT_CLOSE && event.window.windowID == SDL_GetWindowID(window))
                done = true;
        }

        // Start the Dear ImGui frame
        ImGui_ImplOpenGL3_NewFrame();
        ImGui_ImplSDL2_NewFrame(window);
        ImGui::NewFrame();

        {
            ImVec2 mainDisplay = io.DisplaySize;
            memoryProcessesWindow("== Memory and Processes ==",
                                  ImVec2((mainDisplay.x / 2) - 20, (mainDisplay.y / 2) + 30),
                                  ImVec2((mainDisplay.x / 2) + 10, 10));
            // --------------------------------------
            systemWindow("== System ==",
                         ImVec2((mainDisplay.x / 2) - 10, (mainDisplay.y / 2) + 30),
                         ImVec2(10, 10));
            // --------------------------------------
            networkWindow("== Network ==",
                          ImVec2(mainDisplay.x - 20, (mainDisplay.y / 2) - 60),
                          ImVec2(10, (mainDisplay.y / 2) + 50));
        }

        // Rendering
        ImGui::Render();
        glViewport(0, 0, (int)io.DisplaySize.x, (int)io.DisplaySize.y);
        glClearColor(clear_color.x, clear_color.y, clear_color.z, clear_color.w);
        glClear(GL_COLOR_BUFFER_BIT);
        ImGui_ImplOpenGL3_RenderDrawData(ImGui::GetDrawData());
        SDL_GL_SwapWindow(window);
    }

    // Cleanup
    ImGui_ImplOpenGL3_Shutdown();
    ImGui_ImplSDL2_Shutdown();
    ImGui::DestroyContext();

    SDL_GL_DeleteContext(gl_context);
    SDL_DestroyWindow(window);
    SDL_Quit();

    return 0;
}
