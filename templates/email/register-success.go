package email

func ETRegisterSuccess(userName string) string {
	return `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Mailmeteor</title>
    <!--[if mso]><style type="text/css">body, table, td, a { font-family: Arial, Helvetica, sans-serif !important; }</style><![endif]-->
</head>
<body style="font-family: Helvetica, Arial, sans-serif; margin: 0px; padding: 0px; background-color: #ffffff;">
    <table role="presentation"
        style="width: 100%; border-collapse: collapse; border: 0px; border-spacing: 0px; font-family: Arial, Helvetica, sans-serif; background-color: rgb(239, 239, 239);">
        <tbody>
            <tr>
                <td align="center" style="padding: 1rem 2rem; vertical-align: top; width: 100%;">
                    <table role="presentation"
                        style="max-width: 600px; border-collapse: collapse; border: 0px; border-spacing: 0px; text-align: left;">
                        <tbody>
                            <tr>
                                <td style="padding: 40px 0px 0px;">
                                    <div style="padding: 20px; background-color: rgb(255, 255, 255);">
                                        <div style="color: rgb(0, 0, 0); text-align: left;">
                                            <h1 style="margin: 1rem 0">Welcome to WadahGo</h1>
                                            <p style="padding-bottom: 16px">Hello ` + userName + `,</p>
                                            <p style="padding-bottom: 16px">Thank you for signing up to WadahGo.
                                                We're really happy to have you onboard! Click the
                                                link below to login to your account:</p>
                                            <p style="padding-bottom: 16px"><a href="https://wadahgo.com"
                                                    target="_blank"
                                                    style="padding: 12px 24px; border-radius: 4px; color: #FFF; background: #2B52F5;display: inline-block;margin: 0.5rem 0;">Login
                                                    to your account</a></p>
                                            <p style="padding-bottom: 16px">Best regards,</p><span
                                                style="color: #999">WadahGo Limousine</span></p>
                                        </div>
                                    </div>
                                    <div style="padding-top: 20px; color: rgb(153, 153, 153); text-align: center;">
                                        <p style="padding-bottom: 6px;">©Wadahgo 2024, Wadahgo™ is trademark of Wadah
                                            Hub
                                            All rights reserved.</p>
                                        <p>42-01, Jalan Sri Perkasa 1/3, 81200 Johor Bahru</p>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
        </tbody>
    </table>
</body>
</html>`
}

func ETVerifyEmailOTP() {

}
