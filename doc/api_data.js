define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "/home/vladimir/go/src/mdserver/doc/main.js",
    "groupTitle": "/home/vladimir/go/src/mdserver/doc/main.js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/help",
    "title": "Return help page",
    "name": "Help",
    "group": "Common",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Not Found\n{\n   \"message\":\"Failed to open file\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/common/help.go",
    "groupTitle": "Common"
  },
  {
    "type": "post",
    "url": "/log_out",
    "title": "Log out from the site",
    "name": "LogOut",
    "group": "Login",
    "version": "0.0.0",
    "filename": "./internal/handlers/login/log_out.go",
    "groupTitle": "Login"
  },
  {
    "type": "post",
    "url": "/sign_in",
    "title": "Sign in to the site",
    "name": "SignIn",
    "group": "Login",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"username\":\"Vladimir\",\n   \"password\":\"123\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 400 Bad Request\n{\n   \"message\":\"Incorrect login or password\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to save session\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/login/sign_in.go",
    "groupTitle": "Login"
  },
  {
    "type": "post",
    "url": "/create_directory",
    "title": "Create a directory",
    "name": "CreateDirectoryHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"name\":\"Programming\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to create directory\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/create_directory.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/create_post",
    "title": "Create a post",
    "name": "CreatePostHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"dir_name\":\"Programming\",\n   \"file_name\":\"Golang\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to create post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/create_post.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/delete_directory",
    "title": "Delete a directory",
    "name": "DeleteDirectoryHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"name\":\"Programming\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to delete directory\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/delete_directory.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/delete_post",
    "title": "Delete a post",
    "name": "DeletePostHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"dir_name\":\"Programming\",\n   \"file_name\":\"Golang\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to delete post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/delete_post.go",
    "groupTitle": "post"
  },
  {
    "type": "get",
    "url": "/edit/:dir/:file",
    "title": "Get editing post page",
    "name": "EditPostHandler",
    "group": "post",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to get post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/edit_post.go",
    "groupTitle": "post"
  },
  {
    "type": "get",
    "url": "/images/:image",
    "title": "Get image",
    "name": "GetImageHandler",
    "group": "post",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 404 Not Found\n{\n   \"message\":\"Page not found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/get_image.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/edit/:dir/:file/image_upload",
    "title": "Upload the image to temporary files",
    "name": "ImageUploadHandler",
    "group": "post",
    "success": {
      "examples": [
        {
          "title": "Success response example:",
          "content": "    HTTP/1.1 200 OK\n{\n   \"image\":\"e8d4d5cd-d975-4d19-9551-3e29d674ecc3\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 404 Not Found\n{\n   \"message\":\"Image not found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/image_upload.go",
    "groupTitle": "post"
  },
  {
    "type": "get",
    "url": "/:dir/:file",
    "title": "Get post page",
    "name": "PostHandler",
    "group": "post",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to get post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/post.go",
    "groupTitle": "post"
  },
  {
    "type": "get",
    "url": "/",
    "title": "Get index page",
    "name": "PostsHandler",
    "group": "post",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 400 Bad Request\n{\n   \"message\":\"Failed to prepare html\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/posts.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/preview",
    "title": "Get a preview post",
    "name": "PreviewPostHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"text\":\"# Header\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to write response\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/preview_post.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/rename_directory",
    "title": "Rename a directory",
    "name": "RenameDirectoryHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"old_name\":\"Python\",\n   \"new_name\":\"Golang\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to rename directory\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/rename_directory.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/rename_post",
    "title": "Rename a post",
    "name": "RenamePostHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"dir_name\":\"Golang\",\n   \"old_file_name\":\"Patterns\",\n   \"new_file_name\":\"Microservice patterns\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to rename post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/rename_post.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/save_post",
    "title": "Save a post",
    "name": "SavePostHandler",
    "group": "post",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"dir_name\":\"Golang\",\n   \"file_name\":\"Microservice patterns\",\n   \"text\":\"# API gateway\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Error writing the file\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/post/save_post.go",
    "groupTitle": "post"
  },
  {
    "type": "post",
    "url": "/delete_share_link",
    "title": "Delete a share link",
    "name": "DeleteShareLinkHandler",
    "group": "share",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"link\":\"https://vxvxn.ddns.net/share/vxvxn/57afcf14-0b2b-4a1e-be73-d16fbd5df032\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 404 Not Found\n{\n   \"message\":\"Share link not found\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/share/delete_share_link.go",
    "groupTitle": "share"
  },
  {
    "type": "post",
    "url": "/share/:username/:id",
    "title": "Get a share link",
    "name": "GetSharePostHandler",
    "group": "share",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 404 Not Found\n{\n   \"message\":\"Failed to get post\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/share/share.go",
    "groupTitle": "share"
  },
  {
    "type": "get",
    "url": "/share_links",
    "title": "Get list of sharing link",
    "name": "ShareLinksHandler",
    "group": "share",
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Failed to get share link\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/share/share_links.go",
    "groupTitle": "share"
  },
  {
    "type": "post",
    "url": "/share_link",
    "title": "Create a share link",
    "name": "SharePostHandler",
    "group": "share",
    "parameter": {
      "examples": [
        {
          "title": "Request example:",
          "content": "{\n   \"dir_name\":\"Programming\",\n   \"file_name\":\"Golang\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success response example:",
          "content": "    HTTP/1.1 200 OK\n{\n   \"link\":\"https://vxvxn.ddns.net/share/vxvxn/57afcf14-0b2b-4a1e-be73-d16fbd5df032\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "    HTTP/1.1 500 Internal Server Error\n{\n   \"message\":\"Can't create share link to mongo\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/handlers/share/share_post.go",
    "groupTitle": "share"
  }
] });
