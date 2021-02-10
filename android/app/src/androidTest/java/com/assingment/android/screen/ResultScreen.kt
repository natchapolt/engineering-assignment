package com.assingment.android.screen

import com.agoda.kakao.screen.Screen
import com.agoda.kakao.text.KButton
import com.agoda.kakao.text.KTextView
import com.assingment.android.R

class ResultScreen : Screen<ResultScreen>() {
    val name = KTextView { withId(R.id.name) }
    val score = KTextView { withId(R.id.score) }
    val homeButton = KButton { withId(R.id.ctaButton) }
}
