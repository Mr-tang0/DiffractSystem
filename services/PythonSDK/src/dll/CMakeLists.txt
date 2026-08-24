cmake_minimum_required(VERSION 3.15)
project(image_proc C CXX)

set(CMAKE_CXX_STANDARD 17)

# ---- Python 路径 (Miniconda, 可通过命令行覆盖) ----
set(PYTHON_INCLUDE_DIR "D:/Miniconda/Include" CACHE PATH "Python 头文件目录")
set(PYTHON_LIB_DIR "D:/Miniconda/libs" CACHE PATH "Python 库目录")
set(PYTHON_DLL "D:/Miniconda/python313.dll" CACHE FILEPATH "Python DLL 路径")
set(PYTHON3_DLL "D:/Miniconda/python3.dll" CACHE FILEPATH "python3.dll (stable ABI) 路径")

# ---- 构建 DLL ----
add_library(image_proc SHARED image_proc.cpp)
target_include_directories(image_proc PUBLIC
    ${CMAKE_CURRENT_SOURCE_DIR}
    ${PYTHON_INCLUDE_DIR}
)
target_link_directories(image_proc PUBLIC ${PYTHON_LIB_DIR})
target_link_libraries(image_proc python313)

# Windows: 拷贝 python313.dll 和 python3.dll 到输出目录
if(WIN32)
    add_custom_command(TARGET image_proc POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${PYTHON_DLL}" "$<TARGET_FILE_DIR:image_proc>"
        COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${PYTHON3_DLL}" "$<TARGET_FILE_DIR:image_proc>"
        COMMENT "Copying python313.dll and python3.dll"
    )
endif()

# ---- 测试程序 ----
add_executable(test_image_proc test_main.c)
target_link_libraries(test_image_proc image_proc)
