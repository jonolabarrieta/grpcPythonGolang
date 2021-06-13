import asyncio
import logging
import grpc
import data_pb2
import data_pb2_grpc


async def sendMessage(call):  
    while True: 
        print("send to the server")
        await asyncio.sleep(2) 
        await call.write(data_pb2.RouteNote(location="location", message="cliente"))    
        
async def reedMessage(call):
    while True: 
        response = await call.read()
        print(f"Received from the server {response.message} at {response.location}")

async def guide_route_chat(stub: data_pb2_grpc.RouteGuideStub) -> None:
    call = stub.RouteChat()
    sendTask = asyncio.create_task(sendMessage(call))
    redTask =  asyncio.create_task(reedMessage(call))
    await sendTask
    await redTask

async def main() -> None:
    async with grpc.aio.insecure_channel('localhost:10000') as channel:
        stub = data_pb2_grpc.RouteGuideStub(channel)
        await guide_route_chat(stub)
            
if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.get_event_loop().run_until_complete(main())
