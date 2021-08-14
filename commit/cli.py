"""Console script for commit."""
import os
import sys

import click


@click.group(invoke_without_command=True)
@click.pass_context
def main(ctx):
    """Console script for commit."""
    # click.echo("See click documentation at https://click.palletsprojects.com/")
    if ctx.invoked_subcommand is None:
        click.echo("Call default command new...")
        new.main()
    return 0


def check_msg(msg):
    prefix_list = ["feat", "fix", "docs", "style", "refactor", "test", "chore"]
    error_str = "commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore."
    for prefix in prefix_list:
        if msg.startswith(prefix):
            return True
    click.echo(error_str)
    return sys.exit(1)


@main.command()
def format():
    """format the current repo."""
    click.echo(click.style("This will format the code.", fg="green"))
    args = ["black", "--line-length=120", "--target-version=py38"]
    command = " ".join(args)
    os.system(command)


@main.command()
def new():
    """new commit for repo."""
    os.system("git status")
    click.echo(click.style("This will commit all above.", fg="green"))
    msg = input("commit message: ")
    if check_msg(msg):
        os.system("git add .")
        os.system(f'git commit -m "{msg}"')
        os.system("git push")


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
