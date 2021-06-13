import asyncio
import logging
from typing import AsyncIterable
import grpc
import data_pb2
import data_pb2_grpc
import time


class RouteGuideServicer(data_pb2_grpc.RouteGuideServicer):
    async def RouteChat(self, request_iterator: AsyncIterable[data_pb2.RouteNote],unused_context) -> AsyncIterable[data_pb2.RouteNote]:        
        taskSend = asyncio.create_task(sendMessage(unused_context))
        taskRead = asyncio.create_task(reedMessage(unused_context))
        await taskSend
        await taskRead
    
async def sendMessage(unused_context):
    while True:
        await asyncio.sleep(1)
        print("send To the client")
        await unused_context.write(data_pb2.RouteNote(location="aaa", message="Desde server"))

async def reedMessage(unused_context):
    while True: 
        response = await unused_context.read()
        print(f"Received from the client {response.message} at {response.location}")

async def serve() -> None:
    server = grpc.aio.server()
    data_pb2_grpc.add_RouteGuideServicer_to_server( RouteGuideServicer(), server)
    server.add_insecure_port('[::]:10000')
    await server.start()
    await server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.get_event_loop().run_until_complete(serve())
