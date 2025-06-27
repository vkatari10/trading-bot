#include <stdio.h>
#include <unistd.h>
#include <sys/wait.h>

char ** build() {
    char * arg_list[] = {"./scripts/build.sh", NULL};
    return &arg_list;
} // build

char ** env() {
    char * arg_list[] = {"./scripts/env.sh", NULL};
    return &arg_list;
} // env

char ** mlapi() {
    char * arg_list[] = {"./scripts/mlapi.sh", NULL};
    return &arg_list;
} // mlapi

char ** test() {
    char * arg_list[] = {"./scripts/test.sh", NULL};
    return &arg_list;
} // test

char ** train() {
    char * arg_list[] = {"./scripts/train.sh", NULL};
    return &arg_list;
} // train 

int run_child(char * args) {

    int pid = fork();
    int status;

    if (pid == 0) {
        execvp("bash", args);

        perror("exexvp failed");
        return 1;
    } else if (pid > 0) {
        wait(&status);
        return 0;
    } else {
        perror("fork failed");
        return 1;
    } // if-else
} // run_child


int main(int argc, char *argv[]) {


    while (1) {

        char args[64] = {'0'};

        printf("stratforge> ");
        scanf("%s", &args);

        

        printf("%s", args);

    } // while 

} // main
