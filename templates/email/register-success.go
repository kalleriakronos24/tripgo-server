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

func ETOrderSuccess(
	carName string,
	carSamples string,
	customerName string,
	wadahgoLogo string,
	driverName string,
	plateNumber string,
	carImage string,
	person string,
	luggage string,
	paymentMethod string,
	grandTotal string,
	discount string,
	addr1 string,
	addr2 string,
	uid string,
	orderDate string,
	pickupDate string,
	msg string,
	msg2 string,
	msg3 string,
	redirectLink string,
	target string) string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="en">
<head>
<title></title>
<meta charset="UTF-8" />
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<!--[if !mso]>-->
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<!--<![endif]-->
<meta name="x-apple-disable-message-reformatting" content="" />
<meta content="target-densitydpi=device-dpi" name="viewport" />
<meta content="true" name="HandheldFriendly" />
<meta content="width=device-width" name="viewport" />
<meta name="format-detection" content="telephone=no, date=no, address=no, email=no, url=no" />
<style type="text/css">
table {
border-collapse: separate;
table-layout: fixed;
mso-table-lspace: 0pt;
mso-table-rspace: 0pt
}
table td {
border-collapse: collapse
}
.ExternalClass {
width: 100%
}
.ExternalClass,
.ExternalClass p,
.ExternalClass span,
.ExternalClass font,
.ExternalClass td,
.ExternalClass div {
line-height: 100%
}
body, a, li, p, h1, h2, h3 {
-ms-text-size-adjust: 100%;
-webkit-text-size-adjust: 100%;
}
html {
-webkit-text-size-adjust: none !important
}
body, #innerTable {
-webkit-font-smoothing: antialiased;
-moz-osx-font-smoothing: grayscale
}
#innerTable img+div {
display: none;
display: none !important
}
img {
Margin: 0;
padding: 0;
-ms-interpolation-mode: bicubic
}
h1, h2, h3, p, a {
line-height: inherit;
overflow-wrap: normal;
white-space: normal;
word-break: break-word
}
a {
text-decoration: none
}
h1, h2, h3, p {
min-width: 100%!important;
width: 100%!important;
max-width: 100%!important;
display: inline-block!important;
border: 0;
padding: 0;
margin: 0
}
a[x-apple-data-detectors] {
color: inherit !important;
text-decoration: none !important;
font-size: inherit !important;
font-family: inherit !important;
font-weight: inherit !important;
line-height: inherit !important
}
u + #body a {
color: inherit;
text-decoration: none;
font-size: inherit;
font-family: inherit;
font-weight: inherit;
line-height: inherit;
}
a[href^="mailto"],
a[href^="tel"],
a[href^="sms"] {
color: inherit;
text-decoration: none
}
</style>
<style type="text/css">
@media (min-width: 481px) {
.hd { display: none!important }
}
</style>
<style type="text/css">
@media (max-width: 480px) {
.hm { display: none!important }
}
</style>
<style type="text/css">
@media (max-width: 480px) {
.t196{mso-line-height-alt:0px!important;line-height:0!important;display:none!important}.t197{padding-left:30px!important;padding-bottom:40px!important;padding-right:30px!important}.t199,.t259{width:480px!important}.t27{padding-bottom:20px!important}.t194,.t204,.t24,.t242,.t247,.t255,.t29,.t34,.t39,.t44,.t86{width:420px!important}.t26{line-height:28px!important;font-size:26px!important;letter-spacing:-1.04px!important}.t257{padding:40px 30px!important}.t240{padding-bottom:36px!important}.t236{text-align:center!important}.t207,.t209,.t213,.t215,.t219,.t221,.t225,.t227,.t231,.t233,.t53,.t55,.t75,.t77{display:revert!important}.t137,.t188,.t189{display:block!important}.t211,.t217,.t223,.t229,.t235{vertical-align:top!important;width:44px!important}.t80{text-align:left!important}.t57{vertical-align:middle!important;width:221px!important}.t13,.t17{vertical-align:top!important}.t18{text-align:right!important}.t17{width:80px!important}.t15{padding-bottom:50px!important}.t13{width:370px!important}.t4,.t9{width:345.33px!important}.t49{width:353px!important}.t79{vertical-align:middle!important;width:820px!important}.t185,.t59,.t65,.t71{padding-left:0!important}.t61,.t67,.t73{width:303.7px!important}.t188{text-align:left!important}.t137{mso-line-height-alt:15px!important;line-height:15px!important}.t138,.t187{vertical-align:top!important;display:inline-block!important;width:100%!important;max-width:800px!important}.t135{padding-bottom:15px!important;padding-right:0!important}.t117,.t133,.t167,.t183{width:800px!important}.t103,.t108,.t113,.t122,.t128,.t142,.t148,.t153,.t158,.t163,.t172,.t178,.t92,.t98{width:600px!important}
}
</style>
<!--[if !mso]>-->
<link href="https://fonts.googleapis.com/css2?family=Albert+Sans:wght@400;500;700;800&amp;display=swap" rel="stylesheet" type="text/css" />
<!--<![endif]-->
<!--[if mso]>
<xml>
<o:OfficeDocumentSettings>
<o:AllowPNG/>
<o:PixelsPerInch>96</o:PixelsPerInch>
</o:OfficeDocumentSettings>
</xml>
<![endif]-->
</head>
<body id="body" class="t263" style="min-width:100%;Margin:0px;padding:0px;background-color:#242424;"><div class="t262" style="background-color:#242424;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" border="0" align="center"><tr><td class="t261" style="font-size:0;line-height:0;mso-line-height-rule:exactly;background-color:#242424;" valign="top" align="center">
<!--[if mso]>
<v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="true" stroke="false">
<v:fill color="#242424"/>
</v:background>
<![endif]-->
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" border="0" align="center" id="innerTable"><tr><td><div class="t196" style="mso-line-height-rule:exactly;mso-line-height-alt:45px;line-height:45px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t200" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="600" class="t199" style="background-color:#F8F8F8;width:600px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t199" style="background-color:#F8F8F8;width:600px;">
<!--<![endif]-->
<table class="t198" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t197" style="padding:0 50px 60px 50px;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t25" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t24" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t24" style="width:500px;">
<!--<![endif]-->
<table class="t23" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t22"><div class="t21" style="width:100%;text-align:right;"><div class="t20" style="display:inline-block;"><table class="t19" role="presentation" cellpadding="0" cellspacing="0" align="right" valign="top">
<tr class="t18"><td></td><td class="t13" width="370" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t12" style="width:100%;"><tr><td class="t11" style="padding:35px 0 0 0;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t5" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="370" class="t4" style="width:370px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t4" style="width:370px;">
<!--<![endif]-->
<table class="t3" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t2"><p class="t1" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;"><span class="t0" style="margin:0;Margin:0;font-weight:bold;mso-line-height-rule:exactly;">Booking Information</span></p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t10" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="370" class="t9" style="width:370px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t9" style="width:370px;">
<!--<![endif]-->
<table class="t8" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t7" style="padding:0 0 22px 0;"><p class="t6" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">Date: ` + orderDate + `</p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td><td class="t17" width="130" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t16" style="width:100%;"><tr><td class="t15" style="padding:0 0 60px 0;"><div style="font-size:0px;"><img class="t14" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="130" height="130" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/dc9efab5-885c-4675-a477-6c48a2778462.jpeg"/></div></td></tr></table>
</td>
<td></td></tr>
</table></div></div></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t30" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t29" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t29" style="width:500px;">
<!--<![endif]-->
<table class="t28" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t27" style="padding:0 0 15px 0;"><h1 class="t26" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:26px;font-weight:800;font-style:normal;font-size:24px;text-decoration:none;text-transform:none;letter-spacing:-1.56px;direction:ltr;color:#191919;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">Hello ` + customerName + `,</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t35" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t34" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t34" style="width:500px;">
<!--<![endif]-->
<table class="t33" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t32" style="padding:0 0 22px 0;"><p class="t31" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">` + msg + `</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t40" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t39" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t39" style="width:500px;">
<!--<![endif]-->
<table class="t38" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t37" style="padding:0 0 22px 0;"><p class="t36" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">` + msg2 + `</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t45" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t44" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t44" style="width:500px;">
<!--<![endif]-->
<table class="t43" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t42" style="padding:0 0 22px 0;"><p class="t41" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">` + msg3 + `</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="left">
<table class="t50" role="presentation" cellpadding="0" cellspacing="0" style="Margin-right:auto;"><tr>
<!--[if mso]>
<td width="250" class="t49" style="background-color:#181818;overflow:hidden;width:250px;border-radius:44px 44px 44px 44px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t49" style="background-color:#181818;overflow:hidden;width:250px;border-radius:44px 44px 44px 44px;">
<!--<![endif]-->
<table class="t48" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t47" style="text-align:center;line-height:44px;mso-line-height-rule:exactly;mso-text-raise:10px;"><a href="` + redirectLink + `" target="_blank"><span class="t46" style="display:block;margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:44px;font-weight:800;font-style:normal;font-size:12px;text-decoration:none;text-transform:uppercase;letter-spacing:2.4px;direction:ltr;color:#F8F8F8;text-align:center;mso-line-height-rule:exactly;mso-text-raise:10px;">OPEN WadahGo</span></a></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t51" style="mso-line-height-rule:exactly;mso-line-height-alt:40px;line-height:40px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t87" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t86" style="background-color:#F0F0F0;width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t86" style="background-color:#F0F0F0;width:500px;">
<!--<![endif]-->
<table class="t85" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t84" style="padding:20px 20px 20px 20px;"><div class="t83" style="width:100%;text-align:left;"><div class="t82" style="display:inline-block;"><table class="t81" role="presentation" cellpadding="0" cellspacing="0" align="left" valign="middle">
<tr class="t80"><td></td><td class="t57" width="112.36763" valign="middle">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t56" style="width:100%;"><tr><td class="t53" style="width:10px;" width="10"></td><td class="t54"><div style="font-size:0px;"><img class="t52" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="92.36763236763237" height="120.28125" alt="" src="https://gettransfer.com/common/transport_types/premium_small.png"/></div></td><td class="t55" style="width:10px;" width="10"></td></tr></table>
</td><td class="t79" width="387.63237" valign="middle">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t78" style="width:100%;"><tr><td class="t75" style="width:10px;" width="10"></td><td class="t76"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t62" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="367.6323676323676" class="t61" style="width:367.63px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t61" style="width:367.63px;">
<!--<![endif]-->
<table class="t60" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t59" style="padding:0 0 0 10px;"><h1 class="t58" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:700;font-style:normal;font-size:14px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">` + carName + `</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t63" style="mso-line-height-rule:exactly;mso-line-height-alt:10px;line-height:10px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t68" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="367.6323676323676" class="t67" style="width:367.63px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t67" style="width:367.63px;">
<!--<![endif]-->
<table class="t66" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t65" style="padding:0 0 0 10px;"><h1 class="t64" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">` + carSamples + `</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t70" style="mso-line-height-rule:exactly;mso-line-height-alt:15px;line-height:15px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t74" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="367.6323676323676" class="t73" style="border-top:1px solid #CCCCCC;width:367.63px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t73" style="border-top:1px solid #CCCCCC;width:367.63px;">
<!--<![endif]-->
<table class="t72" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t71" style="padding:15px 0 0 10px;"><h1 class="t69" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">Pax: ` + person + `, Luggage: ` + luggage + `</h1></td></tr></table>
</td></tr></table>
</td></tr></table></td><td class="t77" style="width:10px;" width="10"></td></tr></table>
</td>
<td></td></tr>
</table></div></div></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t88" style="mso-line-height-rule:exactly;mso-line-height-alt:30px;line-height:30px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t195" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t194" style="background-color:#F0F0F0;width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t194" style="background-color:#F0F0F0;width:500px;">
<!--<![endif]-->
<table class="t193" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t192" style="padding:40px 40px 40px 40px;"><div class="t191" style="width:100%;text-align:left;"><div class="t190" style="display:inline-block;"><table class="t189" role="presentation" cellpadding="0" cellspacing="0" align="left" valign="top">
<tr class="t188"><td></td><td class="t138" width="210" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t136" style="width:100%;"><tr><td class="t135" style="padding:0 5px 0 0;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t118" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t117" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t117" style="width:205px;">
<!--<![endif]-->
<table class="t116" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t115"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t93" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t92" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t92" style="width:205px;">
<!--<![endif]-->
<table class="t91" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t90"><h1 class="t89" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:700;font-style:normal;font-size:14px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">PICKUP ADDRESS</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t94" style="mso-line-height-rule:exactly;mso-line-height-alt:10px;line-height:10px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t99" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t98" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t98" style="width:205px;">
<!--<![endif]-->
<table class="t97" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t96"><p class="t95" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;">` + addr1 + `</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t104" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t103" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t103" style="width:205px;">
<!--<![endif]-->
<table class="t102" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t101"><p class="t100" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t109" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t108" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t108" style="width:205px;">
<!--<![endif]-->
<table class="t107" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t106"><p class="t105" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t114" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t113" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t113" style="width:205px;">
<!--<![endif]-->
<table class="t112" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t111"><p class="t110" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t130" style="mso-line-height-rule:exactly;mso-line-height-alt:30px;line-height:30px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t134" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t133" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t133" style="width:205px;">
<!--<![endif]-->
<table class="t132" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t131"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t123" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t122" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t122" style="width:205px;">
<!--<![endif]-->
<table class="t121" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t120"><h1 class="t119" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:700;font-style:normal;font-size:14px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">` + target + `</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t124" style="mso-line-height-rule:exactly;mso-line-height-alt:10px;line-height:10px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t129" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t128" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t128" style="width:205px;">
<!--<![endif]-->
<table class="t127" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t126"><p class="t125" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;">` + driverName + ` - ` + plateNumber + `</p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
<!--[if !mso]>-->
<div class="t137" style="mso-line-height-rule:exactly;font-size:1px;display:none;">&nbsp;&nbsp;</div>
<!--<![endif]-->
</td><td class="t187" width="210" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t186" style="width:100%;"><tr><td class="t185" style="padding:0 0 0 5px;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t168" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t167" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t167" style="width:205px;">
<!--<![endif]-->
<table class="t166" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t165"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t143" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t142" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t142" style="width:205px;">
<!--<![endif]-->
<table class="t141" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t140"><h1 class="t139" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:700;font-style:normal;font-size:14px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">DROP-OFF ADDRESS</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t144" style="mso-line-height-rule:exactly;mso-line-height-alt:10px;line-height:10px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t149" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t148" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t148" style="width:205px;">
<!--<![endif]-->
<table class="t147" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t146"><p class="t145" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;">` + addr2 + `</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t154" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t153" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t153" style="width:205px;">
<!--<![endif]-->
<table class="t152" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t151"><p class="t150" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t159" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t158" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t158" style="width:205px;">
<!--<![endif]-->
<table class="t157" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t156"><p class="t155" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t164" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t163" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t163" style="width:205px;">
<!--<![endif]-->
<table class="t162" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t161"><p class="t160" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;"></p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t180" style="mso-line-height-rule:exactly;mso-line-height-alt:30px;line-height:30px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t184" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t183" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t183" style="width:205px;">
<!--<![endif]-->
<table class="t182" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t181"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t173" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t172" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t172" style="width:205px;">
<!--<![endif]-->
<table class="t171" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t170"><h1 class="t169" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:16px;font-weight:700;font-style:normal;font-size:14px;text-decoration:none;text-transform:uppercase;direction:ltr;color:#1A1A1A;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">PAYMENT METHOD</h1></td></tr></table>
</td></tr></table>
</td></tr><tr><td><div class="t174" style="mso-line-height-rule:exactly;mso-line-height-alt:10px;line-height:10px;font-size:1px;display:block;">&nbsp;&nbsp;</div></td></tr><tr><td align="center">
<table class="t179" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="204.99999999999997" class="t178" style="width:205px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t178" style="width:205px;">
<!--<![endif]-->
<table class="t177" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t176"><p class="t175" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#242424;text-align:left;mso-line-height-rule:exactly;mso-text-raise:3px;">` + grandTotal + ` - ` + paymentMethod + `</p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td>
<td></td></tr>
</table></div></div></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t260" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="600" class="t259" style="background-color:#242424;width:600px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t259" style="background-color:#242424;width:600px;">
<!--<![endif]-->
<table class="t258" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t257" style="padding:48px 50px 48px 50px;"><table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="width:100% !important;"><tr><td align="center">
<table class="t205" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t204" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t204" style="width:500px;">
<!--<![endif]-->
<table class="t203" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t202"><p class="t201" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:800;font-style:normal;font-size:18px;text-decoration:none;text-transform:none;letter-spacing:-0.9px;direction:ltr;color:#757575;text-align:center;mso-line-height-rule:exactly;mso-text-raise:1px;">Want updates through more platforms?</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t243" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t242" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t242" style="width:500px;">
<!--<![endif]-->
<table class="t241" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t240" style="padding:10px 0 44px 0;"><div class="t239" style="width:100%;text-align:center;"><div class="t238" style="display:inline-block;"><table class="t237" role="presentation" cellpadding="0" cellspacing="0" align="center" valign="top">
<tr class="t236"><td></td><td class="t211" width="44" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t210" style="width:100%;"><tr><td class="t207" style="width:10px;" width="10"></td><td class="t208"><div style="font-size:0px;"><img class="t206" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="24" height="24" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/5fce123d-03f4-4d3a-964d-b5df10365f92.png"/></div></td><td class="t209" style="width:10px;" width="10"></td></tr></table>
</td><td class="t217" width="44" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t216" style="width:100%;"><tr><td class="t213" style="width:10px;" width="10"></td><td class="t214"><div style="font-size:0px;"><img class="t212" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="24" height="24" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/143f8f5d-1b7f-4cd8-8e88-005411875965.png"/></div></td><td class="t215" style="width:10px;" width="10"></td></tr></table>
</td><td class="t223" width="44" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t222" style="width:100%;"><tr><td class="t219" style="width:10px;" width="10"></td><td class="t220"><div style="font-size:0px;"><img class="t218" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="24" height="24" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/c2ad3667-160d-4ff7-80d6-2bf39d55ba9e.png"/></div></td><td class="t221" style="width:10px;" width="10"></td></tr></table>
</td><td class="t229" width="44" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t228" style="width:100%;"><tr><td class="t225" style="width:10px;" width="10"></td><td class="t226"><div style="font-size:0px;"><img class="t224" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="24" height="24" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/35e4de16-69d5-4814-8bd5-c4086ea16db2.png"/></div></td><td class="t227" style="width:10px;" width="10"></td></tr></table>
</td><td class="t235" width="44" valign="top">
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" class="t234" style="width:100%;"><tr><td class="t231" style="width:10px;" width="10"></td><td class="t232"><div style="font-size:0px;"><img class="t230" style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;" width="24" height="24" alt="" src="https://982f5a97-1e6f-479d-a78d-71b414029431.b-cdn.net/e/7a84f259-a22b-4ec8-86d4-4bdc132a4f87/d05826a6-ea05-495b-83fb-db7231d25354.png"/></div></td><td class="t233" style="width:10px;" width="10"></td></tr></table>
</td>
<td></td></tr>
</table></div></div></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t248" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t247" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t247" style="width:500px;">
<!--<![endif]-->
<table class="t246" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t245"><p class="t244" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;direction:ltr;color:#888888;text-align:center;mso-line-height-rule:exactly;mso-text-raise:3px;">42-01, Jalan Sri Perkasa 1/3, Taman Tampoi Utama, 81200 Johor Bahru, Johor</p></td></tr></table>
</td></tr></table>
</td></tr><tr><td align="center">
<table class="t256" role="presentation" cellpadding="0" cellspacing="0" style="Margin-left:auto;Margin-right:auto;"><tr>
<!--[if mso]>
<td width="500" class="t255" style="width:500px;">
<![endif]-->
<!--[if !mso]>-->
<td class="t255" style="width:500px;">
<!--<![endif]-->
<table class="t254" role="presentation" cellpadding="0" cellspacing="0" width="100%" style="width:100%;"><tr><td class="t253"><p class="t252" style="margin:0;Margin:0;font-family:Albert Sans,BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif;line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;direction:ltr;color:#888888;text-align:center;mso-line-height-rule:exactly;mso-text-raise:3px;"><a class="t249" href="https://tabular.email" style="margin:0;Margin:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#888888;mso-line-height-rule:exactly;" target="_blank">Unsubscribe</a>&nbsp; •&nbsp; <a class="t250" href="https://tabular.email" style="margin:0;Margin:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#888888;mso-line-height-rule:exactly;" target="_blank">Privacy policy</a>&nbsp; •&nbsp; <a class="t251" href="https://tabular.email" style="margin:0;Margin:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#878787;mso-line-height-rule:exactly;" target="_blank">Contact us</a></p></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table>
</td></tr></table>
</td></tr></table></td></tr></table></div><div class="gmail-fix" style="display: none; white-space: nowrap; font: 15px courier; line-height: 0;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</div></body>
</html>`
}
