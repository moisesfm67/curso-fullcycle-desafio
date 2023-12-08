import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma/prisma.service';
import { CreateAssetDto } from './dto/create-asset.dto';

@Injectable()
export class AssetsService {
  constructor(private prisma: PrismaService) {}

  async create(createAssetDto: CreateAssetDto) {
    return this.prisma.asset.create({ data: createAssetDto });
  }

  async findAll() {
    return this.prisma.asset.findMany();
  }
}
