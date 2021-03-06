import { DiscountServiceClient } from '../../protos/discount_grpc_pb'
import { DiscountRequest, Product, DiscountReply, ProductDiscount } from '../../protos/discount_pb'
import { credentials, ServiceError } from '@grpc/grpc-js';

type Resolve<T> = (value: T) => void;
type Reject = (reason?: unknown) => void;

export async function findProductDiscounts(userId: string, productIds: string[]) {
    const address = process.env.NODE_ENV === "production" ? "discounts-service" : "localhost"
    const client = new DiscountServiceClient(address + ':50051', credentials.createInsecure());

    const productsRequest = productIds.map(id => new Product().setId(id))

    const request = new DiscountRequest();
    request.setUserId(userId);
    request.setProductsList(productsRequest);

    const res = await new Promise((resolve: Resolve<DiscountReply>, reject: Reject): void => {
        client.getDiscount(request, (err: ServiceError | null, response: DiscountReply) => {
            if (err) {
                console.error(err);
            }
            resolve(response);
        });
    });

    const discounts = res ? res.getProductsList().map((p: ProductDiscount) => p.toObject()) : [];

    discounts.forEach((discount: ProductDiscount.AsObject) => {
        console.log('Product Id:', discount.productId, 'Percentage:', discount.percentage, 'Value In Cents:', discount.valueInCents);
    })

    return discounts;
}