"""Console script for commit."""
import os
import sys
import click


@click.group()
def main(args=None):
    """Console script for commit."""
    click.echo("Replace this message by putting your code into "
               "commit.cli.main")
    click.echo("See click documentation at https://click.palletsprojects.com/")
    return 0


def check_msg(msg):
    prefix_list = ['feat', 'fix', 'docs', 'style', 'refactor', 'test', 'chore']
    error_str = 'commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore.'
    for prefix in prefix_list:
        if msg.startswith(prefix):
            return True
    click.echo(error_str)
    return sys.exit(1)


@main.command()
def new():
    """new commit for repo."""
    msg = input("commit message: ")
    if check_msg(msg):
        os.system('git add .')
        os.system(f'git commit -m "{msg}"')
        os.system('git push')


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
