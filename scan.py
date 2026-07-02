import socket
import threading
import time

def scan_port(port):
    try:
        conn = socket.create_connection(("192.168.1.1", port), timeout=1)
        conn.close()
    except:
        pass

def main():
    threads = []
    start = time.time()

    for port in range(10000):
        t = threading.Thread(target=scan_port, args=(port,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    print(f"Python (threads): {time.time() - start:.2f} seconds")

if __name__ == "__main__":
    main()