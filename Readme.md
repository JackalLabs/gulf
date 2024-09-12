# Gulf, for Jackal Media Streaming

Gulf is a simple command line tool that takes in a media file, encodes it to HLS format using FFMPEG, and uploads it to Jackal and creates an IPFS folder to index it.

## Install
```shell
git clone https://github.com/JackalLabs/gulf.git
cd gulf
go mod download
go install
```

## Usage
Example using Jackal testnet nodes.
```shell
# you can also use a `.env` file for these
RPC=https://testnet-rpc.jackalprotocol.com:443
GRPC=jackal-testnet-grpc.polkachu.com:17590
SEED={YOUR SEED HERE}

gulf post {YOUR_FILE_HERE}
```
Running this could take a few minutes, so sit back and relax while it does its thing.

Your output should look something like this once it completes:
```shell
Lift Off! 🚀

You can now view your files at ipfs://QmS24SBLrVFWAtEZUE8bu32rDKVRqkE5hPKGBChG9muTCo
```

From there, you can copy that CID and use it in any HLS compatible streamer. You can even paste it into this demo site: https://themarstonconnell.github.io/jackal-streaming-demo/