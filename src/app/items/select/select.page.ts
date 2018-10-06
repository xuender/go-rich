import { Component, OnInit } from '@angular/core';
import { Item } from '../item';
import { ObjsPlusPage } from '../../api/objs.puls.page';
import { ItemService } from '../item.service';
import { TagService } from '../../tags/tag.service';
import { ModalController, ActionSheetController, NavParams } from '@ionic/angular';
import { find, remove } from 'lodash'

@Component({
  selector: 'app-select',
  templateUrl: './select.page.html',
  styleUrls: ['./select.page.scss'],
})
export class SelectPage extends ObjsPlusPage<Item>{
  itemService: ItemService
  items: Item[] = []
  constructor(
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
    private navParams: NavParams
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
    this.itemService = this.navParams.get('service')
  }
  get tagKey() { return 'tag-C' };
  get service() { return this.itemService }
  set service(s) { }
  get page() { return {} }
  get newObj() { return {} }
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