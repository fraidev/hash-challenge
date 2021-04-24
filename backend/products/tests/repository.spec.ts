import { createConnection, getConnection, getRepository } from "typeorm";
import { Product } from "../src/models/Product";
import { TYPEORM_OPTIONS } from "../src/infrastructure/typeorm-options";

describe('products', () => {
    beforeAll(async () => {

    });

    it('GET /products', async () => {
        const db = await createConnection(TYPEORM_OPTIONS);
        const products = await db.getRepository(Product).find();

        expect(products.length).toBeGreaterThan(0);
    });
});