## run

```bash
pipenv install
pipenv run python3 [file.py]
pipenv shell
```

## virtual environment

### 1. pipenv
https://packaging.python.org/en/latest/tutorials/managing-dependencies/#managing-dependencies

1. install pipenv
```python3
python3 -m pip install --user pipenv
python3 -m pip install --user virtualenv
python3 -m pip install --user virtualenv-clone
pipenv install {packages}
```

2. install package for the project

### 2. pyenv

#### install

```bash
brew install pyenv
brew install pyenv-virtualenv
```

#### virtual env

```bash
pyenv install 3.10.9
pyenv virtualenv 3.10.9 nvim
pyenv activate nvim

pip install neovim
pyenv activate nvim
pip install neovim
```

## Trouble Shoot

pipenv --venv 
to findout folder the project environment

https://github.com/pypa/pipenv/issues/5052

