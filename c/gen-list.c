// vim: set sw=4 ts=4 sts=4:

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <time.h>
#include <stdbool.h>
#include "C:\Users\Jim\repos\cJSON\cjson.h"

#define PATHLEN 256
#define MAX_APPS 16384
#define NAME_LEN 64
#ifndef MAX_PATH
#define MAX_PATH 260
#endif

#define CYAN "\e[34m"
#define RESET "\e[0m"

char scoopRoot[]     = "C:/Users/Jim/scoop";
char bucketsDir[]    = "C:/Users/Jim/scoop/buckets";
char appsDir[]       = "C:/Users/Jim/scoop/apps";
int n_apps           = 0;
int n_apps_installed = 0;
char installed_apps[MAX_APPS][NAME_LEN];

typedef struct app {
	char name[NAME_LEN];
	char desc[PATHLEN];
	char bucket[NAME_LEN];
	char homepage[PATHLEN];
	bool installed;
} app_t;

app_t appList[MAX_APPS];


//
// Funcs
//
void load_apps(char * bn, char *dir);
char * read_file(char *path);
void mark_installed_apps();
void print_app(app_t app, FILE*);
void write_apps(FILE *stream);



int main() {

	clock_t t1, t0 = clock();

	DIR *dp = opendir(bucketsDir);
	if (dp == NULL) {
		perror("opendir");
		exit(1);
	}
	struct dirent *e;

	while ((e = readdir(dp)) != NULL) {
		char *bucketName = e->d_name;
		char jsonDir[PATHLEN];
		if (e->d_name[0] != '.') {
			jsonDir[0] = '\0';
			strcat(jsonDir, bucketsDir);
			strcat(jsonDir, "/");
			strcat(jsonDir, e->d_name);
			strcat(jsonDir, "/");
			strcat(jsonDir, "bucket");
			load_apps(bucketName, jsonDir);
		}
	}

	closedir(dp);

	mark_installed_apps();

	char fname[1024];
	strcpy(fname, "app-list-");
	strcat(fname, getenv("COMPUTERNAME"));
	FILE * fp = fopen(fname, "w");
	if (fp) {
		write_apps(fp);
		fclose(fp);
	}
	t1 = clock();

	fprintf(stderr, "time: %f\n", (double) (t1 - t0) / CLOCKS_PER_SEC);

	return 0;
}

bool is_installed(int idx) {
	for (int i = 0; i < n_apps_installed; i++) {
		if (strcmp(appList[idx].name, installed_apps[i]) == 0)
			return true;
	}
	return false;
}

void mark_installed_apps() {

	DIR *dp = opendir(appsDir);

	if (dp == NULL) {
		fprintf(stderr, "could not open dir: %s\n", appsDir);
		exit(1);
	}
	struct dirent *e;

	while ((e = readdir(dp)) != NULL) {
		if (e->d_name[0] != '.') {
			strcpy(installed_apps[n_apps_installed], e->d_name);
			n_apps_installed++;
		}
	}

	for (int i = 0; i < n_apps; i++) {
		appList[i].installed = is_installed(i);
	}

	if (dp)
		closedir(dp);

}


char * read_file(char *path) {
	//fprintf(stderr, "reading file: %s\n", path);
	FILE *fp = fopen(path, "r");
	char *s;
	if (fp != NULL) {
		fseek(fp, 0, SEEK_END);
		long si = ftell(fp);
		//fprintf(stderr, "file len: %ld\n", si);
		rewind(fp);
		s = malloc(sizeof(char) * si + 1);
		fread(s, 1, si, fp);
		fclose(fp);
		s[si] = '\0';
	} else {
		//fprintf(stderr, "could not open file: %s\n", path);
		return NULL;
	}
	//fprintf(stderr, "json len: %lld\n", strlen(s));
	return s;
}


void load_apps(char * bn, char *dir) {

	DIR *dp = opendir(dir);

	if (dp == NULL) {
		perror("opendir");
		exit(1);
	}

	struct dirent *e;
	char json_path[1024];
	cJSON *desc = NULL;
	cJSON *homepage = NULL;

	while ((e = readdir(dp)) != NULL) {
		if (e->d_name[0] != '.') {
			json_path[0] = '\0';
			strcat(json_path, dir);
			strcat(json_path, "/");
			strcat(json_path, e->d_name);
			char *app_name = strdup(e->d_name);
			char * last = strrchr(app_name, '.');
			*last = '\0';

			strcpy(appList[n_apps].name, app_name);
			strcpy(appList[n_apps].bucket, bn);

			//fprintf(stderr, "%s\n", json_path);
			char *content = read_file(json_path);

			cJSON *json = cJSON_Parse(content);

			if (content) free(content);

			if (json == NULL) {
				fprintf(stderr, "error in cJSON_Parse: %s\n", json_path);
				continue;
			}

			desc = cJSON_GetObjectItem(json, "description");

			if (cJSON_IsString(desc) && (desc->valuestring != NULL))
				strcpy(appList[n_apps].desc, desc->valuestring);

			homepage = cJSON_GetObjectItem(json, "homepage");

			if (cJSON_IsString(homepage) && (homepage->valuestring != NULL))
				strcpy(appList[n_apps].homepage, homepage->valuestring);

			n_apps++;


			cJSON_Delete(json);
		}
	}
}


void print_app(app_t app, FILE* stream)
{
	if (app.installed)
		fprintf(stream, "* | %-15.15s", app.name);
	else
		fprintf(stream, "  | %-15.15s", app.name);

	fprintf(stream, " | b:%-10s | %-120.120s", app.bucket, app.desc);
	fprintf(stream, " |%s\n", app.homepage);
}

void write_apps(FILE *stream)
{
	if (stream == NULL)
		stream = stdout;

	for (int i = 0; i < n_apps; i++)
		print_app(appList[i], stream);
}

