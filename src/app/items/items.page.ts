import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';
import { Observable } from 'rxjs';
import { pull, includes } from 'lodash'

import { Item } from './item';
import { ItemPage } from './item/item.page';
import { ItemService } from './item.service';
import { TagService } from '../tags/tag.service';
import { ObjsPage } from '../api/objs.page';

@Component({
  selector: 'app-items',
  templateUrl: './items.page.html',
  styleUrls: ['./items.page.scss'],
})
export class ItemsPage extends ObjsPage<Item> {
  items$: Observable<Item[]>
  tags: string[] = []
  constructor(
    public itemService: ItemService,
    private tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.items$ = this.itemService.all$
  }

  get service() { return this.itemService }
  get page() { return ItemPage }
  get newObj() { return { name: '', price: 0, extend: {}, tags: [] } }
  get title() { return '商品' }

  search(event: CustomEvent) {
    const txt = event.detail.value
    this.items$ = this.itemService.search(txt)
  }

  async select() {
    const tags = await this.tagService.select('tag-I')
    if (tags) {
      this.tags = tags
      this.items$ = this.itemService.select(this.tags)
    }
  }

  hasTag(tag: string): boolean {
    return includes(this.tags, tag)
  }

  removeTag(tag: string) {
    pull(this.tags, tag)
    this.items$ = this.itemService.select(this.tags)
  }
}
