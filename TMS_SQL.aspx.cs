using System;
using System.Data.SqlClient;
using System.Data;

public partial class pages_TMS_TMS_ListAllReq : System.Web.UI.Page
{
    public static string Mode_List = "";
    public static string Mode_List_Str = "";
    public static string uid = "";
    public static string empcode = "";
    public static string id = "";

    private string getCurrConnectionString()
    {
        return "Data Source=database_server;Initial Catalog=database_name;Persist Security Info=True;User ID=developer;Password=AVvEV4Ovf4saFi7UxJTq";
    }

    protected void Page_Load(object sender, EventArgs e)
    {
        Mode_List = Request.QueryString["sho"];
        if (Mode_List == null)
        {
            Mode_List = "1";
        }

        if (Mode_List == "1")
        {
            Mode_List_Str = "My Request";
        }
        else if (Mode_List == "2")
        {
            Mode_List_Str = "Team Request";
        }
        else if (Mode_List == "3")
        {
            Mode_List_Str = "Request";
        }

        if (!string.IsNullOrEmpty(Session["uname"] as string))
        {
            uid = Session["uname"].ToString();
            empcode = Session["Emp_Code_Login"].ToString();
        }
        else
        {
            System.Threading.Thread.Sleep(3000);
            Server.Transfer("TMS_Login.aspx");
        }

        // get parameter from url.
        id = Request["id"] ?? "";

        try
        {
            if (!string.IsNullOrWhiteSpace(id))
            {
                DataTable dataTable;
                SqlDataAdapter dataAdapter;
                string setQuery = "";
                setQuery += " SELECT * ";
                setQuery += " FROM EMPLOYEE ";
                setQuery += " WHERE 1 = 1 ";
                setQuery += " AND WORK_STATUS = '3' ";
                setQuery += " AND EMP_CODE = '" + id + "' ";
                dataAdapter = new SqlDataAdapter(setQuery, getCurrConnectionString());
                dataTable = new DataTable();
                dataAdapter.Fill(dataTable);
                if (dataTable.Rows.Count > 0)
                {
                    userName.InnerHtml = dataTable.Rows[0]["ADDRESS_TEXT"].ToString() + " " + dataTable.Rows[0]["FIRST_NAME_THAI"].ToString() + " " + dataTable.Rows[0]["LAST_NAME_THAI"].ToString();
                }
                else
                {
                    userName.InnerHtml = "Not found!!";
                }
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex.Message);
        }
    }
}
