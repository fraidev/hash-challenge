import { DiscountServiceClient } from '../protos/discount_grpc_pb'
import { DiscountRequest, Product, DiscountReply } from '../protos/discount_pb'
import { credentials, ServiceError } from '@grpc/grpc-js';

async function main() {
    const client = new DiscountServiceClient('localhost:50051', credentials.createInsecure());

    const product = new Product();
    product.setId("155fb96c-9342-4767-8b04-c85e782cb584")

    const request = new DiscountRequest();
    request.setUserId('a1869b98-167c-43d0-9b15-5c8d29c68890');

    request.setProductsList([product]);

    client.getDiscount(request, function (err: ServiceError | null, response: DiscountReply) {
        if (err) {
            console.error(err);
        }

        const products = response.getProductsList().map(p => p.toObject());

        products.forEach(product => {
            console.log('Product Id:', product.productId, 'Percentage:', product.percentage, 'Value In Cents:', product.valueInCents);
        })
    });
}

(async (): Promise<void> => {
    try {
        await main();
    } catch (err) {
        console.error(err);
    }
})();
