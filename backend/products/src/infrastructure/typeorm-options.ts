import { ConnectionOptions } from "typeorm";
import { Product } from "../models/Product";

export const TYPEORM_OPTIONS: ConnectionOptions = {
    type: "postgres",
    host: "localhost",
    port: 5432,
    username: "postgres",
    password: "YOUR_PASSWORD_HERE",
    database: "hash",
    entities: [Product]
}