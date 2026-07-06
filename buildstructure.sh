#!/bin/bash

mkdir -p frontend/assets/css
mkdir -p frontend/assets/images
mkdir -p frontend/assets/icons

mkdir -p frontend/js/api
mkdir -p frontend/js/websocket
mkdir -p frontend/js/pages
mkdir -p frontend/js/components
mkdir -p frontend/js/services
mkdir -p frontend/js/utils
mkdir -p frontend/js/templates

touch frontend/index.html
touch frontend/README.md

touch frontend/assets/css/main.css
touch frontend/assets/css/auth.css
touch frontend/assets/css/navbar.css
touch frontend/assets/css/feed.css
touch frontend/assets/css/post.css
touch frontend/assets/css/comments.css
touch frontend/assets/css/chat.css
touch frontend/assets/css/modal.css
touch frontend/assets/css/responsive.css

touch frontend/js/app.js
touch frontend/js/router.js
touch frontend/js/state.js

touch frontend/js/api/auth.js
touch frontend/js/api/posts.js
touch frontend/js/api/comments.js
touch frontend/js/api/users.js
touch frontend/js/api/chat.js

touch frontend/js/websocket/socket.js
touch frontend/js/websocket/messageEvents.js
touch frontend/js/websocket/onlineEvents.js
touch frontend/js/websocket/notificationEvents.js

touch frontend/js/pages/login.js
touch frontend/js/pages/register.js
touch frontend/js/pages/home.js
touch frontend/js/pages/profile.js
touch frontend/js/pages/error.js

touch frontend/js/components/navbar.js
touch frontend/js/components/sidebar.js
touch frontend/js/components/postCard.js
touch frontend/js/components/commentCard.js
touch frontend/js/components/postForm.js
touch frontend/js/components/commentForm.js
touch frontend/js/components/chatList.js
touch frontend/js/components/chatWindow.js
touch frontend/js/components/messageBubble.js
touch frontend/js/components/onlineUsers.js
touch frontend/js/components/modal.js
touch frontend/js/components/notification.js
touch frontend/js/components/loader.js

touch frontend/js/services/authService.js
touch frontend/js/services/postService.js
touch frontend/js/services/commentService.js
touch frontend/js/services/chatService.js
touch frontend/js/services/websocketService.js

touch frontend/js/utils/fetch.js
touch frontend/js/utils/storage.js
touch frontend/js/utils/validator.js
touch frontend/js/utils/throttle.js
touch frontend/js/utils/debounce.js
touch frontend/js/utils/date.js
touch frontend/js/utils/render.js
touch frontend/js/utils/constants.js

touch frontend/js/templates/loginTemplate.js
touch frontend/js/templates/registerTemplate.js
touch frontend/js/templates/homeTemplate.js
touch frontend/js/templates/chatTemplate.js
touch frontend/js/templates/postTemplate.js

echo "✅ Frontend structure created successfully!"