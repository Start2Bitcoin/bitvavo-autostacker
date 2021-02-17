# Bitvavo autostacker

Ever since the unfortunate exit of [Bittr](https://getbittr.com/) people have been unable to stack sats in an automated way using a Dutch platform.

This software makes it easy to autostack bitcoin on the [Bitvavo](https://bitvavo.com) platform.
The software is completely self-hosted using only Github!

# What does this do
This app checks every hour (or whenever you want) if you have euro's available in your Bitvavo account.
If you don't nothing happens. If you do, the app uses every euro in your account to market buy Bitcoin.
This way, you can setup a reccuring payment from your bank account. As soon as the money arrives at your Bitvavo account, Bitcoin
will be bought with it. 
The app does not (at the moment) withdraw bitcoin to your personal wallet.

The app is hosted completely on Github, and uses Github Actions to check your Bitvavo account every x (default 60) minutes.
Be aware that Github Actions Cron Jobs are sometimes unreliable and may not trigger everytime, and they take some time to trigger for the first time.

# Step by step

- Create a [Bitvavo](https://bitvavo.com) account.
- Make sure your account is fully verified:
    - Upload your ID
    - Verify your phone number
    - Deposit some euro's
    - Add 2FA to your account
- Create a Github account
- Fork this repository (that's the button at the top right of this page)
- In your own version of this repository, go to Settings -> Secrets
- On your Bitvavo account, create an API key that can _View information_ and _Trade_, but *not* withdraw. (This is a safety measure. This software does not (yet) support withdrawals)
    - On Github, click `New repository secret`, add the environment secret with name: `API_KEY` and value -> your Bitvavo API *key* 
    - On Github, click `New repository secret`, add with name: `API_SECRET` and value -> your Bitvavo API *secret*
- Now, Github should run the `Action` every 60 minutes. If you want to check if it has run, check out the `Actions` tab of the repository. Be aware that it may take some time (longer than 60 minutes) to trigger for the first time, and that Github
Actions is not completely reliable in scheduling jobs and may sometimes skip one. But since we check every hour this should not be a problem.
- Make sure to have some euro's in your Bitvavo account to test it out.
