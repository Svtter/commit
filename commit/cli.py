"""Console script for commit."""
import sys
import click


@click.group()
def main(args=None):
    """Console script for commit."""
    click.echo("Replace this message by putting your code into "
               "commit.cli.main")
    click.echo("See click documentation at https://click.palletsprojects.com/")
    return 0


@main.command()
def new():
    """new commit for repo."""
    pass


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
