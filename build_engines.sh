CRPTIFY_DIR="lib/cryptify"
INSTALLER="/env/bin/pyinstaller"

${CRPTIFY_DIR}${INSTALLER} --name=ee1 --onefile ${CRPTIFY_DIR}/encrypt/encryptor.py
${CRPTIFY_DIR}${INSTALLER} --name=ee2 --onefile ${CRPTIFY_DIR}/decrypt/decryptor.py

mv dist/* build/bin/