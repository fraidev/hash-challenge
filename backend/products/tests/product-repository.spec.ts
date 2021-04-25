import { createConnection } from "typeorm";
import { Product } from "../src/models/Product";
import { TYPEORM_OPTIONS } from "../src/infrastructure/typeorm-options";

describe('product repository', () => {
    it('find all products', async () => {
        const db = await createConnection(TYPEORM_OPTIONS);

        const products = await db.getRepository(Product).find();

        await db.close();

        expect(products.length).toBe(4);
        expect(products[0].id).toBe("f0b19f39-b0a8-4b39-a42e-6d83f5cbd4aa")
        expect(products[0].price_in_cents).toBe(470000)
        expect(products[0].title).toBe("Playsation 5")
        expect(products[0].description).toBe("Sony video game")


        expect(products[1].id).toBe("827bd2cc-9537-4330-ad2b-40885344b71a")
        expect(products[1].price_in_cents).toBe(500000)
        expect(products[1].title).toBe("Xbox One X")
        expect(products[1].description).toBe("Microsoft video game")
        
        expect(products[2].id).toBe("8bf94cd1-98db-44dd-b82a-ce7af028b677")
        expect(products[2].price_in_cents).toBe(200000)
        expect(products[2].title).toBe("Switch")
        expect(products[2].description).toBe("Nintendo video game")

        expect(products[3].id).toBe("3d2d1916-e11c-4b87-a514-97b082632b0b")
        expect(products[3].price_in_cents).toBe(500000)
        expect(products[3].title).toBe("PC")
        expect(products[3].description).toBe("A Personal Computer")
    });
});