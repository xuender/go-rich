import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';

import { Item } from './item';
import { Paging } from '../api/paging';
import { ItemPage } from './item/item.page';
import { ItemService } from './item.service';
import { TagService } from '../tags/tag.service';
import { ObjsPlusPage } from '../api/objs.puls.page';

@Component({
  selector: 'app-items',
  templateUrl: './items.page.html',
  styleUrls: ['./items.page.scss'],
})
export class ItemsPage extends ObjsPlusPage<Item> {
  tags: string[] = []
  paging: Paging<Item> = { data: [], total: 0 }
  constructor(
    public itemService: ItemService,
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
  }

  get tagKey() { return 'tag-I' };
  get service() { return this.itemService }
  get page() { return ItemPage }
  get newObj() { return { name: '', price: 0, extend: {}, tags: [] } }
  get title() { return '商品' }
}
