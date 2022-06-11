package emails

import (
	"fmt"
	"sister-backend/utils"
	"time"
)

func GenerateProductSyncNotificationEmail(product_count int) utils.MailMessage {
	msg := utils.NewMailMessage()

	msg.SetSubject("Product Sync Notification")
	msg.SetBody(fmt.Sprintf("<p>%d products were successfully synced at %s </p>", product_count, time.Now().Format(time.RFC1123)))
	return *msg
}
