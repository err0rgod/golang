#include <iostream>
#include <thread>
#include <vector>
#include <chrono>
#include <winsock2.h>
#include <ws2tcpip.h>

#pragma comment(lib, "ws2_32.lib")

void scan_port(int port) {
    SOCKET sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == INVALID_SOCKET) return;

    // 1 second timeout
    DWORD timeout = 1000;
    setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, (char*)&timeout, sizeof(timeout));
    setsockopt(sock, SOL_SOCKET, SO_SNDTIMEO, (char*)&timeout, sizeof(timeout));

    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    inet_pton(AF_INET, "192.168.1.1", &addr.sin_addr);

    connect(sock, (struct sockaddr*)&addr, sizeof(addr));
    closesocket(sock);
}

int main() {
    // winsock init — required on windows
    WSADATA wsaData;
    WSAStartup(MAKEWORD(2, 2), &wsaData);

    std::vector<std::thread> threads;
    auto start = std::chrono::high_resolution_clock::now();

    for (int i = 1; i < 1000; i++) {
        threads.emplace_back(scan_port, i);
    }

    for (auto& t : threads) {
        t.join();
    }

    auto end = std::chrono::high_resolution_clock::now();
    double seconds = std::chrono::duration<double>(end - start).count();
    std::cout << "C++ (threads): " << seconds << " seconds" << std::endl;

    WSACleanup();
    return 0;
}