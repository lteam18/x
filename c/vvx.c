#include "stdlib.h"
#include "stdio.h"
#include "string.h"

char** slice(int argc, char** argv, int start){
    char** ptr = malloc( ( argc - start) * sizeof(char*) );
    for (int i=0; i<argc - start; ++i) {
        ptr[i] = argv[start + i];
    }
    return ptr;
}

char* join(char* sep, int argv_len, char* argv[]){
    int sep_len = strlen(sep);
    
    int len = sep_len * (argv_len - 1);
    for (int i=0; i<argv_len; ++i) {
        len += strlen(argv[i]);
    }

    char* ret = malloc(len+1);
    ret[0] = 0;
    for (int i=0; i<argv_len; ++i) {
        strcat(ret, argv[i]);
        if (i != argv_len - 1) strcat(ret, sep);
    }
    return ret;
}

FILE* exists(char* filepath){
    return fopen(filepath, "rb");
}

char* getHomeDirectory(){
    return getenv("HOME");
}

char* locateFilepath(char* filepath){
    if ( NULL != exists(filepath) ) return filepath;
    if ('@' == filepath[0]) {
        char* tmp_arr[] = { getenv("HOME"), "/.vvsh/APP/", filepath+1 };
        filepath = join("", 3, tmp_arr);
        if ( NULL != exists(filepath) ) return filepath;
    }
    return NULL;
}

#define BUFFER_SIZE 1024
char BUFFER[BUFFER_SIZE];

int catFile(char* filepath){
    FILE* fp = fopen(filepath, "r");
    if ( NULL == fp ) {
        printf("[ERROR] Fail to open file %s\n", filepath);
        return 1;
    }
    while ( NULL != fgets(BUFFER, BUFFER_SIZE, fp) ){
        fputs(BUFFER, stdout);
    }
    fclose(fp);
    return 0;
}

int catFiles(int argc, char* filepath_list[]){
    for (int i=0; i<argc; ++i) {
        char* filepath = locateFilepath(filepath_list[i]);
        if ( NULL == filepath) {
            // TODO: From now on, you should cat the files in vvsh-go
            char* cmd[] = { "vvx-go", join(" ", argc-i, slice(argc, filepath_list, i)) };
            system(join(" ", 2, cmd));
            return 0;
        }

        catFile(filepath);
    }
    return 0;
}

int runScript(char* cmd, int argc, char* argv[]) {
    char** ptr = malloc( ( argc + 1) * sizeof(char*) );
    ptr[0] = cmd;
    for (int i=0; i<argc; ++i) {
        ptr[i+1] = argv[i];
    }
    system(join(" ", argc + 1, ptr));
    return 0;
}

void vvxInGo(char* cmd, int argc, char** argv){
    char* s[] = { "vvx-go", cmd, join(" ", argc, argv) };
    char* run_cmd = join(" ", 3, s);
    system(run_cmd);
}

int invokeScript(char* cmd, int argc, char** argv) {
    if (0 == argc) {
        system(cmd);
        return 0;
    }

    char* filepath = locateFilepath(argv[0]);
    if (NULL == filepath) {
        vvxInGo(cmd, argc, argv);
    } else {
        argv[0] = filepath;
        runScript(cmd, argc, argv);
    }
    return 0;
}

int main(int argc, char *argv[]){
    if (1 == argc) {
        system("vvx-go version");
        return 0;
    }

    const char* sub = argv[1];

    if (strcmp(sub, "hi") == 0) {
        printf("hi");
        return 0;
    }

    if (strcmp(sub, "version") == 0) {
        system("vvx-go version");
        return 0;
    }

    // if command name is run, js, node
    if (strcmp(sub, "bash") == 0) {
        invokeScript("bash", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "fish") == 0) {
        invokeScript("fish", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "perl") == 0) {
        invokeScript("perl", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "python") == 0) {
        invokeScript("python", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "cat") == 0) {
        catFiles(argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "java") == 0) {
        invokeScript("java -jar", argc - 2, slice(argc, argv, 2));
        // Invoke jar
        return 0;
    }

    if (strcmp(sub, "js") == 0) {
        invokeScript("vvsh", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "vvsh") == 0) {
        invokeScript("vvsh", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    if (strcmp(sub, "node") == 0) {
        invokeScript("node", argc - 2, slice(argc, argv, 2));
        return 0;
    }

    // printf("\nF: %s, %s\n", sub, join(" ", argc - 2, slice(argc, argv, 2)));
    vvxInGo((char*)sub, argc - 2, slice(argc, argv, 2));
    return 0;
}
