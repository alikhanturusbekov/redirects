# redirects

Technical desicions and their reasons:
- I have chosen Fiber GO framework to code faster and easier (as this is a small project and does not require special features). 
Note: With larger projects I would prefer using combinations of GO modules over frameworks.

- I have used GORM to show my experience with ORM; however, I understand that using ORM highly affects performance in negative way. 
As for smaller project like this, using GORM is adequate.

- I have used code structure recommended by GO best practices and utilized architecture for GO Fiber projects.

# .env-template

DB_USER=alikhan
DB_PASSWORD=alikhan
DB_NAME=redirects

# some details

- I simulated the behaviour of admin using headers {"User-Role"} to manage access to different endpoints. So you need to pass "User-Role" in headers to access admin endpoints

Thanks for your time :)
