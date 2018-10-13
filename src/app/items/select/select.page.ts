import { find, remove } from 'lodash'
import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';

import { Item } from '../item';
import { ItemService } from '../item.service';
import { TagService } from '../../tags/tag.service';
import { ObjsPlusPage } from '../../api/objs.puls.page';
import { ItemPage } from '../item/item.page';

@Component({
  selector: 'app-select',
  templateUrl: './select.page.html',
  styleUrls: ['./select.page.scss'],
})
export class SelectPage extends ObjsPlusPage<Item>{
  items: Item[] = []
  constructor(
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
    public itemService: ItemService,
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
  }
  get tagKey() { return 'tag-I' };
  get service() { return this.itemService }
  set service(s) { }
  get page() { return ItemPage }
  get newObj() { return { name: '', price: 0, extend: {}, tags: [] } }
  get title() { return '商品' }

  trigger(item: Item) {
    if (find(this.items, i => i.id == item.id)) {
      remove(this.items, i => i.id == item.id)
    } else {
      this.items.push(item)
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.modalCtrl.dismiss(this.items)
  }
}
