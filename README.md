***P99 Auctions (V2.0)***

As many were disappointed, the original P99Auctions site was retired after an individual bombarded the site with too much traffic forcing the guy operating the site into submission; this project aims to rebuild the great tool that existed with a hope to extend its functionality further.

All code here is open source and may be forked by anyone to modify/extend but please note that any changes made and submitted to this master branch will be subject to code review before the pull request is merged to master, I will be looking for some trusted admin staff to assist me with this moderation step.

***Important***

Please note that right now all of the code is amalgamated in this auction-parser-client repository.   I will very soon be migrating a bunch of the code which scrapes the P99 Wiki for item data (that doesn't exist in our DB) to the auction-parser-server repo (which doesn't exist yet).  This is fairly sloppy on my end and I hope to get this cleared up soon!

***Dependencies***

If you download this now, the only dependency required is my small stringutil library /alexmk92/stringutil it should install automatically when you run `go install` on your machine, please note that once your run `go install` you need to run `auction-parser-client` from your terminal, if you cd into the src directory and try to execute `go run main.go` you'll get a bunch of linker errors, this is because the src hasn't been compiled and although the resources exist within the same package they are not linked until bundled with `go install` from the root directory.

If you have any questions you can contact me at alexander.sims92@gmail.com I'd be happy to answer any requests!