using System;
using System.Data;
using System.Data.SqlClient;

public class SqlInjectionVulnerable
{
    private static string GetConnectionString()
    {
        return "Data Source=database_server;Initial Catalog=database_name;Persist Security Info=True;User ID=developer;Password=developer";
    }

    public static void GetUserData(string userId)
    {
        // Unsafe query concatenation leading to SQL injection vulnerability
        string query = "SELECT * FROM Users WHERE UserID = '" + userId + "'";

        using (SqlConnection connection = new SqlConnection(GetConnectionString()))
        {
            using (SqlCommand command = new SqlCommand(query, connection))
            {
                connection.Open();

                using (SqlDataReader reader = command.ExecuteReader())
                {
                    while (reader.Read())
                    {
                        Console.WriteLine($"UserID: {reader["UserID"]}");
                        Console.WriteLine($"UserName: {reader["UserName"]}");
                        // Add more fields as needed
                    }
                }
            }
        }
    }

    public static void Main(string[] args)
    {
        // Simulated user input with SQL injection payload
        string userId = "example_user_id' OR '1'='1"; // Example of SQL injection payload

        // Fetch user data
        GetUserData(userId);
    }
}
