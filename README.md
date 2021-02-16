# Bitvavo autostacker

Ever since the unfortunate exit of [Bittr](https://getbittr.com/) people have been unable to stack sats in an automated way using a Dutch platform.

This software makes it easy to autostack bitcoin on the [Bitvavo](https://bitvavo.com) platform.
The software is completely self-hosted using Github and [Vercel](https://vercel.com)

# What does this do
This app checks (by default) every 60 minutes if you have euro's available in your Bitvavo account.
If you don't nothing happens. If you do, the app uses every euro in your account to market buy Bitcoin.
This way, you can setup a reccuring payment from your bank account. As soon as the money arrives at your Bitvavo account, Bitcoin
will be bought with it. You will receive e-mail confirmations from Bitvavo notifying you that the buy has taken place.

The app should be hosted by Vercel, and we use a Github Actions Cron job to poll Bitvavo every x (60) minutes.
All this is hosted for free.

# Step by step

- Create a [Bitvavo](https://bitvavo.com) account.
- Make sure your account is fully verified:
    - Upload your ID
    - Verify your phone number
    - Deposit some euro's
    - Add 2FA to your account
- Create a Github account
- Fork this repository (that's the button at the top right of this page)
- Create a [Vercel](https://vercel.com) account. You can use your Github account to do this. Vercel is needed to host this software.
- Create a new project on Vercel from this repository that you have just forked.
- You will be prompted to add environment variables. Keep this window open.
- Create an API key that can _View information_ and _Trade_, but *not* withdraw. (This is a safety measure. This software does not (yet) support withdrawals)
    - In Vercel, add the environment variables `CONFIG_APIKEY` -> your Bitvavo API *key*
    - In Vercel, add the environment variables `CONFIG_APISECRET` -> your Bitvavo API *secret*
    - Deploy the Vercel app
- In your own repository, update the `.github/workflows/workflow.yaml` file.
- Replace `YOUR_VERCEL_URL` with the url of your app that you get from Vercel.
- (Optional) Replace `60` with how often you want the app to check (in minutes) that you have euro's available in your Bitvavo account.
- Commit and push the changes