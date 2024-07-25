import sqlite3

# Vulnerable function that takes user input to query a SQLite database
def get_user(username):
    conn = sqlite3.connect('users.db')
    cursor = conn.cursor()
    
    # Vulnerable SQL query construction
    query = "SELECT * FROM users WHERE username = '" + username + "'"
    
    cursor.execute(query)
    user = cursor.fetchone()
    
    conn.close()
    return user

# Example usage (vulnerable)
username = input("Enter username: ")
user = get_user(username)

if user:
    print("User found:", user)
else:
    print("User not found")
