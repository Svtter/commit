from invoke import task
import os


@task
def clean(c):
    os.remove('commit.exe')


@task
def test(c):
    """run all test in this package.

    Args:
        c ([type]): [description]
    """
    c.run('go test ./...')
