package com.assingment.android.view.quiz

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import androidx.fragment.app.Fragment
import androidx.fragment.app.activityViewModels
import androidx.fragment.app.viewModels
import androidx.navigation.fragment.findNavController
import androidx.viewpager2.widget.ViewPager2
import com.assingment.android.R
import com.assingment.android.view.ShareViewModel

class QuizFragment: Fragment() {

    private val layoutId: Int = R.layout.fragment_quiz
    private val quizViewModel : QuizViewModel by viewModels()
    private val model: ShareViewModel by activityViewModels()

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(layoutId, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val button = view.findViewById<Button>(R.id.ctaButton)
        val pager = view.findViewById<ViewPager2>(R.id.pager)
        val next = view.findViewById<Button>(R.id.nextButton)
        val back = view.findViewById<Button>(R.id.backButton)

        button.setOnClickListener {
            findNavController().navigate(R.id.action_quizFragment_to_scoreFragment)
        }

        quizViewModel.quizQuestions.observe(viewLifecycleOwner, { it ->
            pager.adapter = QuizAdapter(it, model)
        })

        pager.registerOnPageChangeCallback(object : ViewPager2.OnPageChangeCallback() {
            override fun onPageSelected(position: Int) {
                super.onPageSelected(position)
                quizViewModel.page = position
            }
        })

        next.setOnClickListener {
            pager.setCurrentItem(quizViewModel.page + 1, true)
        }

        back.setOnClickListener {
            pager.setCurrentItem(quizViewModel.page - 1, true)
        }
    }
}
