<%@ Page Language="C#" EnableEventValidation="false" AutoEventWireup="true" CodeFile="TMS_SQL.aspx.cs" Inherits="pages_TMS_TMS_ListAllReq" %>

<%@ Import Namespace="System" %>
<%@ Import Namespace="System.Collections.Generic" %>
<%@ Import Namespace="System.Web" %>
<%@ Import Namespace="System.Web.UI" %>
<%@ Import Namespace="System.Web.UI.WebControls" %>
<%@ Import Namespace="System.Data" %>
<%@ Import Namespace="System.Data.SqlClient" %>
<%@ Import Namespace="OfficeOpenXml" %>
<%@ Import Namespace="OfficeOpenXml.Style" %>
<%@ Import Namespace="System.Drawing" %>

<!DOCTYPE html>
<html>
<head>
    <title>POC Test - SQL Injection</title>
</head>
<body>
    <p id="userName" runat="server">Running...</p>
</body>
</html>

<script src="../../plugins/jQuery/jQuery-2.1.3.min.js"></script>

<script type="text/javascript">
    const constBaseUrl = 'TMS_SQL.aspx';

    const setFormBody = (bodyData, initBodyKey = '') => {
        if (initBodyKey) return JSON.stringify({ [initBodyKey]: JSON.stringify(bodyData) });
        else return JSON.stringify({ 'FormBody': JSON.stringify(bodyData) });
    }

    const httpPost = (methodUrl, bodyData, newBaseUrl = '', initBodyKey = '') => {
        const requestUrl = newBaseUrl ? `${newBaseUrl}/${methodUrl}` : `${constBaseUrl}/${methodUrl}`;
        return new Promise(function (resolve, reject) {
            $.ajax({
                type: "POST",
                url: requestUrl,
                data: setFormBody(bodyData, initBodyKey),
                async: false,
                contentType: "application/json; charset=utf-8",
                dataType: "json",
                success: function (response) { resolve(response.d); },
                failure: function (error) { reject(error.d); }
            });
        });
    }

    function clientCallAPI() {
        // define the api endpoint
        const apiurl = 'https://appserv1.bigth.com:3185/qas/e-workflow/master/business'; // replace with your api endpoint

        // make the get request
        fetch(apiurl, {
            method: 'get',
            headers: {
                'content-type': 'application/json',
                'x-auth-bigth': 'ULoocUynIxZ9xTolfrWoUrAdl5YJnkr5'
            }
        })
            .then(response => {
                if (!response.ok) {
                    throw new error('network response was not ok');
                }
                return response.json(); // parse the json response
            })
            .then(data => {
                console.log('======================= Start - Client Side =======================');
                console.log({ Status: 'success', Data: data }); // handle the response data
                console.log('======================= End - Client Side =======================');
            })
            .catch(error => {
                console.log('======================= Start - Client Side =======================');
                console.error({ Status: 'error', Data: error }); // handle any errors
                console.log('======================= End - Client Side =======================');
            });
    }

    async function Todo() {
        clientCallAPI();
        await serverCallAPI();
    }

    async function serverCallAPI() {
        const bodyData = {};
        const httpResult = await httpPost('GetBusiness', bodyData);
        const httpResponse = JSON.parse(httpResult)

        console.log('======================= Start - Server Side =======================');
        console.log({ Status: httpResponse.Status, Data: JSON.parse(httpResponse.Data) });
        console.log('======================= End - Server Side =======================');
    }

    //Todo();
</script>
