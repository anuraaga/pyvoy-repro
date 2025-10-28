```bash
git clone https://github.com/anuraaga/pyvoy-repro.git
cd pyvoy-repro
uv sync --python 3.13 --managed-python
uv run pyvoy kitchensink
# Separate terminal
go run ./main.go
```
