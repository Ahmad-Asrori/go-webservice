package util

import (
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "fans.askingalexandria26@gmail.com"
const CONFIG_PASSWORD = "tidaktahu"

func SendEmail(to, username, subject, activationCode string) error {
	mailer := gomail.NewMessage()
	template := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:o=\"urn:schemas-microsoft-com:office:office\">\n<head>\n    <!--[if gte mso 9]>\n    <xml>\n        <o:OfficeDocumentSettings>\n            <o:AllowPNG/>\n            <o:PixelsPerInch>96</o:PixelsPerInch>\n        </o:OfficeDocumentSettings>\n    </xml>\n    <![endif]-->\n    <meta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, maximum-scale=1\" />\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\" />\n    <meta name=\"format-detection\" content=\"date=no\" />\n    <meta name=\"format-detection\" content=\"address=no\" />\n    <meta name=\"format-detection\" content=\"telephone=no\" />\n    <meta name=\"x-apple-disable-message-reformatting\" />\n    <!--[if !mso]><!-->\n    <link href=\"https://fonts.googleapis.com/css?family=Muli:400,400i,700,700i\" rel=\"stylesheet\" />\n    <!--<![endif]-->\n    <title>Email Template</title>\n    <!--[if gte mso 9]>\n    <style type=\"text/css\" media=\"all\">\n        sup { font-size: 100% !important; }\n    </style>\n    <![endif]-->\n\n\n    <style type=\"text/css\" media=\"screen\">\n        /* Linked Styles */\n        body { padding:0 !important; margin:0 !important; display:block !important; min-width:100% !important; width:100% !important; background:#001736; -webkit-text-size-adjust:none }\n        a { color:#66c7ff; text-decoration:none }\n        p { padding:0 !important; margin:0 !important }\n        img { -ms-interpolation-mode: bicubic; /* Allow smoother rendering of resized image in Internet Explorer */ }\n        .mcnPreviewText { display: none !important; }\n\n\n        /* Mobile styles */\n        @media only screen and (max-device-width: 480px), only screen and (max-width: 480px) {\n            .mobile-shell { width: 100% !important; min-width: 100% !important; }\n            .bg { background-size: 100% auto !important; -webkit-background-size: 100% auto !important; }\n\n            .text-header,\n            .m-center { text-align: center !important; }\n\n            .center { margin: 0 auto !important; }\n            .container { padding: 20px 10px !important }\n\n            .td { width: 100% !important; min-width: 100% !important; }\n\n            .m-br-15 { height: 15px !important; }\n            .p30-15 { padding: 30px 15px !important; }\n\n            .m-td,\n            .m-hide { display: none !important; width: 0 !important; height: 0 !important; font-size: 0 !important; line-height: 0 !important; min-height: 0 !important; }\n\n            .m-block { display: block !important; }\n\n            .fluid-img img { width: 100% !important; max-width: 100% !important; height: auto !important; }\n\n            .column,\n            .column-top,\n            .column-empty,\n            .column-empty2,\n            .column-dir-top { float: left !important; width: 100% !important; display: block !important; }\n\n            .column-empty { padding-bottom: 10px !important; }\n            .column-empty2 { padding-bottom: 30px !important; }\n\n            .content-spacing { width: 15px !important; }\n        }\n    </style>\n</head>\n<body class=\"body\" style=\"padding:0 !important; margin:0 !important; display:block !important; min-width:100% !important; width:100% !important; background:#001736; -webkit-text-size-adjust:none;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" bgcolor=\"#001736\">\n    <tr>\n        <td align=\"center\" valign=\"top\">\n            <table width=\"650\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" class=\"mobile-shell\">\n                <tr>\n                    <td class=\"td container\" style=\"width:650px; min-width:650px; font-size:0pt; line-height:0pt; margin:0; font-weight:normal; padding:55px 0px;\">\n                        <!-- Header -->\n                        <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                            <tr>\n                                <td class=\"p30-15\" style=\"padding: 0px 30px 30px 30px;\">\n                                    <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                        <tr>\n                                            <th class=\"column-top\" width=\"145\" style=\"font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;\">\n                                                <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                                    <tr>\n                                                        <td class=\"img m-center\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"><img src=\"images/logo.jpg\" width=\"131\" height=\"38\" border=\"0\" alt=\"\" /></td>\n                                                    </tr>\n                                                </table>\n                                            </th>\n                                            <th class=\"column-empty\" width=\"1\" style=\"font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;\"></th>\n                                            <th class=\"column\" style=\"font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal;\">\n                                                <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                                    <tr>\n                                                        <td class=\"text-header\" style=\"color:#475c77; font-family:'Muli', Arial,sans-serif; font-size:12px; line-height:16px; text-align:right;\"><a href=\"#\" target=\"_blank\" class=\"link2\" style=\"color:#475c77; text-decoration:none;\"><span class=\"link2\" style=\"color:#475c77; text-decoration:none;\">Open in your browser</span></a></td>\n                                                    </tr>\n                                                </table>\n                                            </th>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                        </table>\n                        <!-- END Header -->\n\n                        <!-- Intro -->\n                        <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                            <tr>\n                                <td style=\"padding-bottom: 10px;\">\n                                    <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                        <tr>\n                                            <td class=\"tbrr p30-15\" style=\"padding: 60px 30px; border-radius:26px 26px 0px 0px;\" bgcolor=\"#12325c\">\n                                                <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                                    <tr>\n                                                        <td class=\"h1 pb25\" style=\"color:#ffffff; font-family:'Muli', Arial,sans-serif; font-size:40px; line-height:46px; text-align:center; padding-bottom:25px;\">Welcome, " + username + "</td>\n                                                    </tr>\n                                                    <tr>\n                                                        <td class=\"text-center pb25\" style=\"color:#c1cddc; font-family:'Muli', Arial,sans-serif; font-size:16px; line-height:30px; text-align:center; padding-bottom:25px;\">Welcome to Krama Apps <span class=\"m-hide\"><br /></span>This is Your Activation Account.</td>\n                                                    </tr>\n                                                    <!-- Button -->\n                                                    <tr>\n                                                        <td align=\"center\">\n                                                            <table class=\"center\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" style=\"text-align:center;\">\n                                                                <tr>\n                                                                    <td class=\"pink-button text-button\" style=\"background:#ff6666; color:#c1cddc; font-family:'Muli', Arial,sans-serif; font-size:14px; line-height:18px; padding:12px 30px; text-align:center; border-radius:0px 22px 22px 22px; font-weight:bold;\"><p class=\"link-white\" style=\"color:#ffffff; text-decoration:none;\"><span class=\"link-white\" style=\"color:#ffffff; text-decoration:none;\">" + activationCode +"</span></p></td>\n                                                                </tr>\n                                                            </table>\n                                                        </td>\n                                                    </tr>\n                                                    <!-- END Button -->\n                                                </table>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                        </table>\n                        <!-- END Intro -->\n\n                        <!-- Footer -->\n                        <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                            <tr>\n                                <td class=\"p30-15 bbrr\" style=\"padding: 50px 30px; border-radius:0px 0px 26px 26px;\" bgcolor=\"#0e264b\">\n                                    <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                        <tr>\n                                            <td align=\"center\" style=\"padding-bottom: 30px;\">\n                                                <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                                                    <tr>\n                                                        <td class=\"img\" width=\"55\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"><a href=\"#\" target=\"_blank\"><img src=\"images/ico_facebook.jpg\" width=\"38\" height=\"38\" border=\"0\" alt=\"\" /></a></td>\n                                                        <td class=\"img\" width=\"55\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"><a href=\"#\" target=\"_blank\"><img src=\"images/ico_twitter.jpg\" width=\"38\" height=\"38\" border=\"0\" alt=\"\" /></a></td>\n                                                        <td class=\"img\" width=\"55\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"><a href=\"#\" target=\"_blank\"><img src=\"images/ico_instagram.jpg\" width=\"38\" height=\"38\" border=\"0\" alt=\"\" /></a></td>\n                                                        <td class=\"img\" width=\"38\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"><a href=\"#\" target=\"_blank\"><img src=\"images/ico_linkedin.jpg\" width=\"38\" height=\"38\" border=\"0\" alt=\"\" /></a></td>\n                                                    </tr>\n                                                </table>\n                                            </td>\n                                        </tr>\n                                        <tr>\n                                            <td class=\"text-footer1 pb10\" style=\"color:#c1cddc; font-family:'Muli', Arial,sans-serif; font-size:16px; line-height:20px; text-align:center; padding-bottom:10px;\">Krama Apps - Testing Application</td>\n                                        </tr>\n                                        <tr>\n                                            <td class=\"text-footer2\" style=\"color:#8297b3; font-family:'Muli', Arial,sans-serif; font-size:12px; line-height:26px; text-align:center;\">Jakarta, Indonesia<br />&copy; 2020 | krama Apps</td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                        </table>\n                        <!-- END Footer -->\n                    </td>\n                </tr>\n            </table>\n        </td>\n    </tr>\n</table>\n</body>\n</html>\n"

	mailer.SetHeader("From", CONFIG_EMAIL)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", template)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_EMAIL,
		CONFIG_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}