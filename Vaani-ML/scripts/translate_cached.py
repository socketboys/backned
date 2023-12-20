

REPO_ID = "facebook/nllb-200-distilled-600M"
# FILENAME = "sklearn_model.joblib"s

from huggingface_hub import snapshot_download

snapshot_download(repo_id=REPO_ID)