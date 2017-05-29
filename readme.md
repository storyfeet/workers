Workers is a simple project to get a fixed number of workers running a specific set of func() , no return for safety.

You can see some accompanying insight in [this YouTube video](https://youtu.be/CUG1vfnO3zI).

-- note a small change in the code from the video, thanks to Devarsh Shah, Who pointed out, the workers can range on res.Main instead of checking inside the for loop.

so

    for {
        f := <-res.Main
        if f != nil {
            procDone <- true
            return
        }
        f()
    }

becomes

    for f := range res.Main {
        f()
    }
    procDone <- true

