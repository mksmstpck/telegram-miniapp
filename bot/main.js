import {Telegraf, Markup} from "telegraf"

const token = "7375982299:AAEZ7yRQswUFGbCG6uVxLi27zmVHzRS87pE"

const bot = new Telegraf(token)

bot.command("start", (ctx) => {
    ctx.reply("Шо ти єбать?")
})

bot.launch()