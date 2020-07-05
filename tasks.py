from invoke import task
import os


@task
def clean(c):
    os.remove('commit.exe')
