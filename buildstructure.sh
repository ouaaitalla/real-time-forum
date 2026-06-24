#!/bin/bash

mkdir -p backend/{db,handlers,models,websocket}
touch backend/main.go

mkdir -p frontend/{css,js}
touch frontend/index.html
touch frontend/js/{app.js,auth.js,forum.js,chat.js,websocket.js}

mkdir -p database
touch database/forum.db

echo "Project structure created successfully!"