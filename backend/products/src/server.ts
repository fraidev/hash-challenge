import express from 'express'
import { createConnection } from "typeorm";
import { ProductDiscount } from '../protos/discount_pb';
import { TYPEORM_OPTIONS } from './infrastructure/typeorm-options';
import { Product } from "./models/Product";
import { findProductDiscounts } from './services/discount-service';

const app = express()
const port = 8080

createConnection(TYPEORM_OPTIONS).then(connection => {
    const productRepository = connection.getRepository(Product);

    app.get('/product', async (req, res) => {
        const userId = req.headers['user-id'] as string;
        if (!userId) {
            res.status(400).send('USER_ID header is required')
        }

        const products = await productRepository.find();

        const discounts = await findProductDiscounts(userId, products.map(p => p.id))

        const result = products.map(p => {
            const discountFound = discounts.find((d: ProductDiscount.AsObject) => d.productId === p.id);
            if (discountFound){
                p.setDiscount(discountFound.percentage, discountFound.valueInCents)
            }
            return p
        })

        console.log(result)

        res.json(result)
    })

    app.listen(port, () => {
        console.log(`Product app listening at http://localhost:${port}/product`)
    })
})