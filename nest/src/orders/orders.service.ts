import { Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { PrismaService } from 'src/prisma/prisma/prisma.service';

@Injectable()
export class OrdersService {
  constructor(private prisma: PrismaService) {}

  async create(createOrderDto: CreateOrderDto) {
    return this.prisma.order.create({
      data: { ...createOrderDto, status: 'PENDING' },
    });
  }

  async findAll() {
    return this.prisma.order.findMany();
  }
}
