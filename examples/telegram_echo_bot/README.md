# Telegram Echo Bot Example

این مثال نشان می‌دهد چگونه می‌توانید با استفاده از کتابخانه‌ی [telego](https://github.com/mymmrac/telego) یک ربات ساده بسازید که به پیام‌های کاربران پاسخ دهد.

## راه‌اندازی

1. یک ربات جدید از طریق [BotFather](https://t.me/BotFather) ایجاد کنید و توکن دسترسی را دریافت نمایید.
2. متغیر محیطی `TELEGRAM_BOT_TOKEN` را با مقدار توکن تنظیم کنید:

   ```bash
   export TELEGRAM_BOT_TOKEN="your-telegram-token"
   ```

3. برنامه را اجرا کنید:

   ```bash
   go run ./examples/telegram_echo_bot
   ```

پس از اجرا، هر پیامی را که به ربات ارسال کنید با یک پاسخ ساده دریافت خواهید کرد و دستورات `/start` و `/help` نیز پیام راهنما ارسال می‌کنند.
